package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"code.miracl.com/maas/maas/src/lib/gomiracl"
	"code.miracl.com/maas/maas/src/lib/gomiracl/wrap"
)

func createSession(httpClient *http.Client, userID string) (*sessionResponse, error) {
	sessionRequest := &struct {
		ProjectID string `json:"projectId"`
		UserID    string `json:"userId"`
	}{
		options.projectID,
		userID,
	}

	sessionResp, err := makeRequest(
		httpClient,
		options.apiURL+"/rps/v2/session",
		http.MethodPost,
		sessionRequest,
		header{Key: "Content-Type", Value: "application/json"})
	if err != nil {
		return nil, err
	}

	var createSessionResponse *sessionResponse

	if err := json.Unmarshal(sessionResp, &createSessionResponse); err != nil {
		return nil, err
	}

	return createSessionResponse, nil
}

func register(httpClient *http.Client, userID, deviceName string, pin int, accessID string) (i identity, err error) {
	// Call to /verification endpoint.
	verifyURL, err := verificationRequest(httpClient, userID, deviceName, accessID)
	if err != nil {
		return identity{}, err
	}

	verificationCode, err := getVerificationCode(verifyURL)
	if err != nil {
		return identity{}, err
	}

	activationToken, err := verificationConfirmation(httpClient, userID, verificationCode)
	if err != nil {
		return identity{}, err
	}

	id, err := newIdentity(httpClient, userID, deviceName, accessID, activationToken, pin)
	if err != nil {
		return identity{}, err
	}

	return id, nil
}

func newIdentity(httpClient *http.Client, userID, deviceName, accessID, activationToken string, pin int) (i identity, err error) {
	// Call to /rps/v2/user endpoint.
	regResponse, err := registerRequest(httpClient, userID, deviceName, accessID, activationToken)
	if err != nil {
		return identity{}, err
	}

	// Call to /signature endpoint.
	sigResponse, err := signatureRequest(httpClient, regResponse.MPinID, regResponse.RegOTT)
	if err != nil {
		return identity{}, err
	}

	// Call to /dta/ID endpoint.
	csResponse, err := clientSecretRequest(httpClient, sigResponse.CS2URL)
	if err != nil {
		return identity{}, err
	}

	// Combine both client secrets.
	Q, err := wrap.RecombineG1BN254CX(hex2bytes(sigResponse.ClientSecretShare), hex2bytes(csResponse.ClientSecret))
	if err != nil {
		return identity{}, err
	}

	// First extract pin from the combine client secret, in order to get the token.
	CS, err := wrap.ExtractPINBN254CX(int(gomiracl.SHA256), hex2bytes(regResponse.MPinID), pin, Q)
	if err != nil {
		return identity{}, err
	}

	return identity{
		MPinID: hex2bytes(regResponse.MPinID),
		Token:  CS,
		DTAs:   sigResponse.DTAs,
	}, nil
}

func verificationRequest(httpClient *http.Client, userID, deviceName, accessID string) (string, error) {
	clientIDAndSecret := options.clientID + ":" + options.clientSecret
	authHeaderValue := "Basic " + b64.StdEncoding.EncodeToString([]byte(clientIDAndSecret))

	payload := struct {
		ProjectID     string `json:"projectId"`
		UserID        string `json:"userId"`
		DeviceName    string `json:"deviceName"`
		AccessID      string `json:"accessId"`
		Delivery      string `json:"delivery"`
		Authorization string `json:"-"`
	}{
		options.projectID,
		userID,
		deviceName,
		accessID,
		"no",
		authHeaderValue,
	}

	resp, err := makeRequest(
		httpClient,
		options.apiURL+"/verification",
		"POST",
		payload,
		header{Key: "Authorization", Value: authHeaderValue},
	)
	if err != nil {
		return "", err
	}

	var verificationResponse *verificationURLResponse

	if err := json.Unmarshal(resp, &verificationResponse); err != nil {
		return "", err
	}

	return verificationResponse.VerificationURL, nil
}

func registerRequest(httpClient *http.Client, userID, deviceName, accessID, activateCode string) (*registerResponse, error) {
	payload := struct {
		UserID       string `json:"userId"`
		DeviceName   string `json:"deviceName"`
		WID          string `json:"wid"`
		ActivateCode string `json:"activateCode"`
	}{
		DeviceName:   deviceName,
		UserID:       userID,
		WID:          accessID,
		ActivateCode: activateCode,
	}

	resp, err := makeRequest(
		httpClient,
		options.apiURL+"/rps/v2/user",
		"PUT",
		payload,
		header{Key: "X-MIRACL-CID", Value: "mcl"},
	)
	if err != nil {
		return nil, err
	}

	var registerResponse *registerResponse
	if err := json.Unmarshal(resp, &registerResponse); err != nil {
		return nil, err
	}

	return registerResponse, nil
}

var errInvalidSignatureResponse = errors.New("invalid signature response")

func signatureRequest(httpClient *http.Client, mpinID, regOTT string) (*signatureResponse, error) {
	resp, err := makeRequest(
		httpClient,
		fmt.Sprintf(options.apiURL+"/rps/v2/signature/%v?regOTT=%v", mpinID, regOTT),
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}

	var sigResponse *signatureResponse

	if err := json.Unmarshal(resp, &sigResponse); err != nil {
		return nil, err
	}

	if !(sigResponse.CS2URL != "" && sigResponse.ClientSecretShare != "" &&
		sigResponse.Curve != "" && sigResponse.DTAs != "") {
		return nil, errInvalidSignatureResponse
	}

	return sigResponse, nil
}

func clientSecretRequest(httpClient *http.Client, cs2url string) (*clientSecretResponse, error) {
	resp, err := makeRequest(
		httpClient,
		cs2url,
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}

	var csResponse *clientSecretResponse

	if err := json.Unmarshal(resp, &csResponse); err != nil {
		return nil, err
	}

	return csResponse, nil
}

func getVerificationCode(verifyURL string) (string, error) {
	parsedURL, err := url.Parse(verifyURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse verification url: %w", err)
	}

	return parsedURL.Query().Get("code"), nil
}

func verificationConfirmation(httpClient *http.Client, userID, code string) (string, error) {
	payload := &confirmationRequest{
		UserID: userID,
		Code:   code,
	}

	resp, err := makeRequest(
		httpClient,
		options.apiURL+"/verification/confirmation",
		"POST",
		payload,
	)
	if err != nil {
		return "", fmt.Errorf("error creating verification confirmation request: %w", err)
	}

	var res *confirmationResponse

	if err := json.Unmarshal(resp, &res); err != nil {
		return "", err
	}

	return res.ActivateToken, nil
}
