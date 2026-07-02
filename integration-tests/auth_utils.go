package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"code.miracl.com/maas/maas/src/lib/gomiracl"
	"code.miracl.com/maas/maas/src/lib/gomiracl/curve"
	"code.miracl.com/maas/maas/src/lib/gomiracl/mpin"
)

var errAuthenticate = errors.New("error making the authenticate request")

func authenticate(httpClient *http.Client, identity *identity, pin int, code string) error {
	token := mpin.NewClientToken(identity.Token, curve.BN254CX)

	// Client Pass 1.
	pass1Proof, err := token.Pass1(identity.dvsMPinID(), pin, gomiracl.SHA256, newRand())
	if err != nil {
		return fmt.Errorf("error getting client 1: %w", err)
	}

	// Call to /rps/v2/pass1 endpoint for Server Pass 1.
	p1Response, err := pass1Request(httpClient, identity.MPinID, identity.PublicKey, pass1Proof.U, pass1Proof.UT, identity.DTAs, "jwt")
	if err != nil {
		return fmt.Errorf("error making the pass 1 request: %w", err)
	}

	y, err := hex.DecodeString(p1Response.Y)
	if err != nil {
		return fmt.Errorf("error decoding pass1 response: %w", err)
	}

	// Client Pass 2.
	pass2Proof, err := token.Pass2(pass1Proof.X, y, pass1Proof.SEC)
	if err != nil {
		return fmt.Errorf("error getting client 2: %w", err)
	}

	// Call to /rps/v2/pass2 endpoint for Server Pass 2.
	p2Response, err := pass2Request(httpClient, identity.MPinID, pass2Proof.V, code)
	if err != nil {
		return fmt.Errorf("error making the pass 2 request: %w", err)
	}

	// Call to /rps/v2/authenticate endpoint.
	authResponse, err := authenticateRequest(httpClient, p2Response.AuthOTT)
	if err != nil {
		return fmt.Errorf("error making the authenticate request: %w", err)
	}

	if authResponse.Status != http.StatusOK {
		return fmt.Errorf("%w, status: %v", errAuthenticate, authResponse.Status)
	}

	return nil
}

func pass1Request(httpClient *http.Client, mpinID, pubKey, u, ut []byte, dtas string, scope ...string) (p1Response *pass1Response, err error) {
	payload := &struct {
		DTAs      string   `json:"dtas"`
		MPinID    string   `json:"mpin_id"`
		Pass      int      `json:"pass"`
		PublicKey string   `json:"publicKey"`
		Scope     []string `json:"scope"`
		U         string   `json:"U"`
		UT        string   `json:"UT"`
	}{
		DTAs:      dtas,
		MPinID:    hex.EncodeToString(mpinID),
		Pass:      1,
		PublicKey: hex.EncodeToString(pubKey),
		Scope:     scope,
		U:         hex.EncodeToString(u),
		UT:        hex.EncodeToString(ut),
	}

	resp, err := makeRequest(
		httpClient,
		options.projectDomain+"/rps/v2/pass1",
		http.MethodPost,
		payload,
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp, &p1Response); err != nil {
		return nil, err
	}

	return p1Response, nil
}

func pass2Request(httpClient *http.Client, mpinID, proof []byte, wid string) (p2Response *pass2Response, err error) {
	payload := &struct {
		MPinID string `json:"mpin_id"`
		V      string `json:"V"`
		WID    string `json:"WID"`
	}{
		MPinID: hex.EncodeToString(mpinID),
		V:      hex.EncodeToString(proof),
		WID:    wid,
	}

	resp, err := makeRequest(
		httpClient,
		options.projectDomain+"/rps/v2/pass2",
		http.MethodPost,
		payload,
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp, &p2Response); err != nil {
		return nil, err
	}

	return p2Response, nil
}

func authenticateRequest(httpClient *http.Client, authOTT string) (authResponse *authenticateResponse, err error) {
	payload := struct {
		AuthOTT string `json:"authOTT"`
	}{
		authOTT,
	}

	resp, err := makeRequest(
		httpClient,
		options.projectDomain+"/rps/v2/authenticate",
		http.MethodPost,
		payload,
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp, &authResponse); err != nil {
		return nil, err
	}

	return authResponse, nil
}

func accessRequest(httpClient *http.Client, webOTT string) (accessResponse *accessResponse, err error) {
	payload := struct {
		WebOTT string `json:"webOTT"`
	}{
		webOTT,
	}

	resp, err := makeRequest(
		httpClient,
		options.projectDomain+"/rps/v2/access",
		http.MethodPost,
		payload,
	)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resp, &accessResponse); err != nil {
		return nil, err
	}

	return accessResponse, nil
}
