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

	httpClient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	sessionResponse, err := createSession(httpClient, userID)
	if err != nil {
		t.Fatalf("failed to create session: %v", err.Error())
	}

	qrURL, err := url.Parse(sessionResponse.QRURL)
	if err != nil {
		t.Fatalf("failed to parse session QR URL: %v", err.Error())
	}

	accessID := qrURL.Fragment

	identity, err := register(httpClient, userID, deviceName, pin, accessID)
	if err != nil {
		t.Fatalf("Error registering: %v", err)
	}

	err = authenticate(httpClient, identity, pin, accessID)
	if err != nil {
		t.Fatalf("Error authenticating: %v", err)
	}

	accessResponse, err := accessRequest(httpClient, sessionResponse.WebOTT)
	if err != nil {
		t.Fatalf("error making the access request: %v", err.Error())
	}

	if accessResponse.UserID != userID {
		t.Fatal("UserID mismatch")
	}
}

func TestValidateSignature(t *testing.T) {
	if options.skipModifyTests {
		t.Skip("skip TestValidateSignature")
	}

	if options.proxyHost == "" || options.proxyPort == "" || options.sampleName == "" {
		t.Fatal("sample-name, proxy-host or proxy-port are missing")
	}

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

	sessionResponse, err := createSession(httpClient, userID)
	if err != nil {
		t.Fatalf("failed to create session: %v", err.Error())
	}

	qrURL, err := url.Parse(sessionResponse.QRURL)
	if err != nil {
		t.Fatalf("failed to parse QR URL: %v", err.Error())
	}

	accessID := qrURL.Fragment

	identity, err := register(httpClient, userID, deviceName, pin, accessID)
	if err != nil {
		t.Fatalf("Error registering: %v", err)
	}

	err = authenticate(httpClient, identity, pin, accessID)
	if err != nil {
		t.Fatalf("Error authenticating: %v", err)
	}

	// Call to /rps/v2/access endpoint.
	accessResponse, err := accessRequest(httpClient, sessionResponse.WebOTT)
	if err != nil {
		t.Fatalf("error making the access request: %v", err.Error())
	}

	if accessResponse.UserID != userID {
		t.Fatal("UserID mismatch")
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
