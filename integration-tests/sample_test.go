package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	if options.skipModifyTests {
		t.Skip("skip TestValidateSignature")
	}

	if options.proxyHost == "" || options.proxyPort == "" || options.sampleName == "" {
		t.Fatal("sample-name, proxy-host or proxy-port are missing")
	}

	// All samples generate a new state and redirect us to an /authorize URL, if we're not logged in.
	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	client := newSampleClient(options.sampleURL, httpClient)
	client.restart(options.restarterHost, options.restarterPort, options.sampleName)

	m := newModifier(modifySignatureHandler, options.proxyHost, options.proxyPort, httpClient)
	m.start()

	defer m.stop()

	name := uuid.New().String()
	userID := fmt.Sprintf("%v@example.com", name)
	deviceName := "The device of " + name
	pin := randPIN()

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
}

func modifySignatureHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("ERROR is: %#v", err.Error())))
		return
	}

	defer r.Body.Close()

	const jwksUri = "https://api.mpin.io:443/oidc/certs"

	originalRequestURL := r.Header.Get("X-Forwarded-Host")
	if originalRequestURL == jwksUri {
		keys := struct {
			Keys []map[string]string `json:"keys"`
		}{}

		if err = json.Unmarshal(body, &keys); err != nil {
			w.Write([]byte(fmt.Sprintf("/oidc/certs keys decoding error: %#v", err.Error())))
			return
		}

		keys.Keys[0]["n"] = "invalid"
		keys.Keys[1]["n"] = "invalid"

		body, err = json.Marshal(keys)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("keys encoding error: %#v", err.Error())))
			return
		}
	}

	for k, v := range r.Header {
		if k != "X-Forwarded-Host" {
			for _, hv := range v {
				w.Header().Set(k, hv)
			}
		}
	}

	w.Write(body)
}
