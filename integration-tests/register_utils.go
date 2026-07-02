package main

import (
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"code.miracl.com/maas/maas/src/lib/gomiracl"
	"code.miracl.com/maas/maas/src/lib/gomiracl/curve"
	"code.miracl.com/maas/maas/src/lib/gomiracl/mpin"
)

func createSession(httpClient *http.Client, projectID, userID string) (*sessionResponse, error) {
	sessionRequest := &struct {
		ProjectID string `json:"projectId"`
		UserID    string `json:"userId"`
	}{
		projectID,
		userID,
	}

	sessionResp, err := makeRequest(
		httpClient,
		options.projectDomain+"/rps/v2/session",
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

func register(httpClient *http.Client, projectID, userID, deviceName string, pin int) (i *identity, err error) {
	// Call to /verification endpoint.
	verifyURL, err := verificationRequest(httpClient, userID, deviceName, projectID)
	if err != nil {
		return nil, err
	}

	verificationCode, err := getVerificationCode(verifyURL)
	if err != nil {
		return nil, err
	}

	// Call to /verification/confirmation endpoint.
	activationToken, err := verificationConfirmation(httpClient, userID, verificationCode)
	if err != nil {
		return nil, err
	}

	id, err := newIdentity(httpClient, userID, deviceName, activationToken, pin)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func getProjectID(httpClient *http.Client) (projectID string, err error) {
	resp, err := makeRequest(
		httpClient,
		options.projectDomain+"/.well-known/project-configuration",
		http.MethodGet,
		nil,
	)
	if err != nil {
		return "", err
	}

	var projectResponse *struct {
		ID string `json:"id"`
	}

	if err := json.Unmarshal(resp, &projectResponse); err != nil {
		return "", err
	}

	return projectResponse.ID, nil
}

func newIdentity(httpClient *http.Client, userID, deviceName, activationToken string, pin int) (i *identity, err error) {
	privateKey, publicKey, err := curve.BN254CX.GetDVSKeyPair(newRand(), nil)
	if err != nil {
		return nil, fmt.Errorf("error generating key pair: %w", err)
	}

	// Call to /registration endpoint.
	regResponse, err := registrationRequest(httpClient, userID, deviceName, hex.EncodeToString(publicKey), activationToken)
	if err != nil {
		return nil, err
	}

	mpinID, err := hex.DecodeString(regResponse.MPinID)
	if err != nil {
		return nil, err
	}

	// Call to all SecretURLs to get client secrets and combine them to create a client token.
	clientToken, err := clientTokenRequests(httpClient, regResponse.SecretURLs, mpinID, publicKey, privateKey, pin)
	if err != nil {
		return nil, err
	}

	return &identity{
		MPinID:     mpinID,
		Token:      clientToken,
		DTAs:       regResponse.DTAs,
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}, nil
}

func verificationRequest(httpClient *http.Client, userID, deviceName, projectID string) (string, error) {
	clientIDAndSecret := options.clientID + ":" + options.clientSecret
	authHeaderValue := "Basic " + b64.StdEncoding.EncodeToString([]byte(clientIDAndSecret))

	payload := struct {
		ProjectID     string `json:"projectId"`
		UserID        string `json:"userId"`
		DeviceName    string `json:"deviceName"`
		Delivery      string `json:"delivery"`
		Authorization string `json:"-"`
	}{
		projectID,
		userID,
		deviceName,
		"no",
		authHeaderValue,
	}

	resp, err := makeRequest(
		httpClient,
		options.projectDomain+"/verification",
		http.MethodPost,
		payload,
		header{Key: "Authorization", Value: authHeaderValue},
	)
	if err != nil {
		return "", err
	}

	var verificationResponse verificationURLResponse

	if err := json.Unmarshal(resp, &verificationResponse); err != nil {
		return "", err
	}

	return verificationResponse.VerificationURL, nil
}

func registrationRequest(httpClient *http.Client, userID, deviceName, publicKey, activationToken string) (*registrationResponse, error) {
	payload := &struct {
		ActivationToken string `json:"activationToken"`
		DeviceName      string `json:"deviceName"`
		PublicKey       string `json:"publicKey"`
		UserID          string `json:"userId"`
	}{
		ActivationToken: activationToken,
		DeviceName:      deviceName,
		PublicKey:       publicKey,
		UserID:          userID,
	}

	resp, err := makeRequest(
		httpClient,
		options.projectDomain+"/registration",
		http.MethodPost,
		payload,
		header{Key: "X-MIRACL-CID", Value: "mcl"},
	)
	if err != nil {
		return nil, err
	}

	var registrationResponse registrationResponse

	if err := json.Unmarshal(resp, &registrationResponse); err != nil {
		return nil, err
	}

	return &registrationResponse, nil
}

func clientTokenRequests(httpClient *http.Client, secretURLs []string, mpinID, publicKey, privateKey []byte, pin int) (token []byte, err error) {
	const secretURLsNumber = 2
	if len(secretURLs) != secretURLsNumber {
		return nil, fmt.Errorf("client secret urls should be exactly 2; received: %v", len(secretURLs))
	}

	css := make([][]byte, 0, secretURLsNumber)

	for _, secretURL := range secretURLs {
		resp, err := makeRequest(
			httpClient,
			secretURL,
			http.MethodGet,
			nil,
		)
		if err != nil {
			return nil, err
		}

		var cssResponse clientSecretResponse

		if err := json.Unmarshal(resp, &cssResponse); err != nil {
			return nil, err
		}

		secret, err := hex.DecodeString(cssResponse.DVSClientSecret)
		if err != nil {
			return nil, err
		}

		css = append(css, secret)
	}

	clientToken, err := mpin.GetDVSClientToken(
		append(mpinID, publicKey...),
		pin,
		gomiracl.SHA256,
		privateKey,
		curve.BN254CX,
		css[0],
		css[1],
	)
	if err != nil {
		return nil, err
	}

	return clientToken.Secret(), nil
}

func getVerificationCode(verifyURL string) (string, error) {
	parsedURL, err := url.Parse(verifyURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse verification url: %w", err)
	}

	return parsedURL.Query().Get("code"), nil
}

func verificationConfirmation(httpClient *http.Client, userID, code string) (string, error) {
	payload := &struct {
		UserID string `json:"userId"`
		Code   string `json:"code"`
	}{
		UserID: userID,
		Code:   code,
	}

	resp, err := makeRequest(
		httpClient,
		options.projectDomain+"/verification/confirmation",
		http.MethodPost,
		payload,
	)
	if err != nil {
		return "", fmt.Errorf("error creating verification confirmation request: %w", err)
	}

	var res confirmationResponse

	if err := json.Unmarshal(resp, &res); err != nil {
		return "", err
	}

	return res.ActivateToken, nil
}
