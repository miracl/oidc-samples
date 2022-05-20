package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestAuth(t *testing.T) {
	name := uuid.New().String()
	userID := fmt.Sprintf("%v@example.com", name)
	deviceName := "The device of " + name
	pin := randPIN()

	// All samples generate a new state and redirect us to an /authorize URL, if we're not logged in.
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	client := newSampleClient(options.sampleURL, httpClient)

	authorizeRequestURL, err := client.authorize()
	if err != nil {
		t.Fatalf("Error making initial request to the home page: %v", err)
	}

	_, err = url.ParseRequestURI(string(authorizeRequestURL))
	if err != nil {
		t.Fatalf("Home page didn't respond with authorize url, error: %v, actual response: %v", err.Error(), string(authorizeRequestURL))
	}

	identity, err := register(httpClient, userID, deviceName, pin, string(authorizeRequestURL))
	if err != nil {
		t.Fatalf("Error registering: %v", err)
	}

	accessResponse, err := authenticate(httpClient, identity, userID, pin, string(authorizeRequestURL))
	if err != nil {
		t.Fatalf("Error authenticating: %v", err)
	}

	// Contains an URL, from which we can fetch the user info.
	err = client.login(accessResponse.RedirectURL)
	if err != nil {
		t.Fatalf("Error logging in: %v", err)
	}

	userInfo, err := client.getUserInfo()
	if err != nil {
		t.Fatalf("Error getting user info: %v", err)
	}

	if userInfo.Email != userID {
		t.Fatal("UserID  mismatch")
	}
}

func TestValidateSignature(t *testing.T) {
	if os.Getenv("MODIFY") != "true" {
		t.Skip("Skipping testing if not in MODIFY environment")
	}

	if options.proxyHost == "" || options.proxyPort == "" || options.modifyHost == "" || options.modifyPort == "" {
		t.Fatal("proxy-host, proxy-port, modify-host and modify-port are required to run this test")
	}

	// All samples generate a new state and redirect us to an /authorize URL, if we're not logged in.
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	payload := struct {
		ModifyURL string `json:"modifyUrl"`
	}{
		ModifyURL: "http://" + options.modifyHost + ":" + options.modifyPort + "/modifySignature", //127.0.0.1:8081/modifySignature",
	}

	prreq, err := newRequest(fmt.Sprintf("http://%v:%v/session", options.proxyHost, options.proxyPort), "POST", payload)
	if err != nil {
		t.Fatalf("Error making request to the proxy: %v", err)
	}

	_, _, err = getResponse(prreq, httpClient)
	if err != nil {
		t.Fatalf("Error when receiving the proxy response: %v", err)
	}

	name := uuid.New().String()
	userID := fmt.Sprintf("%v@example.com", name)
	deviceName := "The device of " + name
	pin := randPIN()

	client := newSampleClient(options.sampleURL, httpClient)

	authorizeRequestURL, err := client.authorize()
	if err != nil {
		t.Fatalf("Error making initial request to the home page: %v", err)
	}

	_, err = url.ParseRequestURI(string(authorizeRequestURL))
	if err != nil {
		t.Fatalf("Home page didn't respond with authorize url, error: %v, actual response: %v", err.Error(), string(authorizeRequestURL))
	}

	identity, err := register(httpClient, userID, deviceName, pin, string(authorizeRequestURL))
	if err != nil {
		t.Fatalf("Error registering: %v", err)
	}

	accessResponse, err := authenticate(httpClient, identity, userID, pin, string(authorizeRequestURL))
	if err != nil {
		t.Fatalf("Error authenticating: %v", err)
	}

	// Contains an URL, from which we can fetch the user info.
	err = client.login(accessResponse.RedirectURL)
	if err == nil {
		t.Fatal("Expected error when the signature is invalid")
	}

	if !errors.Is(err, errInvalidLogin) {
		t.Fatal("Expected invalid login error")
	}

	if !strings.Contains(err.Error(), "Invalid ID token") {
		t.Fatalf("Expected Invalid ID token error message")
	}

	// stop the proxy modifying session
	prreq, err = newRequest(fmt.Sprintf("http://%v:%v/session", options.proxyHost, options.proxyPort), "DELETE", http.NoBody)
	if err != nil {
		t.Fatalf("Error when making stop session request to the proxy: %v", err)
	}

	_, _, err = getResponse(prreq, httpClient)
	if err != nil {
		t.Fatalf("Error on proxy response: %v", err)
	}
}
