package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
)

type app struct {
	html           *template.Template
	stateGenerator func() string
	stateStorage   set
	sessionStorage set
	oauthConfig    *oauth2.Config
	verifier       *oidc.IDTokenVerifier
	provider       *oidc.Provider
}

func newApp(f *flags) (*app, error) {
	provider, err := oidc.NewProvider(context.Background(), f.issuer)
	if err != nil {
		return nil, err
	}

	app := &app{
		stateGenerator: func() string { return strconv.Itoa(rand.Intn(999999)) },
		stateStorage:   newSet(),
		oauthConfig: &oauth2.Config{
			ClientID:     f.clientID,
			ClientSecret: f.clientSecret,
			Endpoint:     provider.Endpoint(),
			RedirectURL:  f.redirectURL,
			Scopes:       []string{oidc.ScopeOpenID, "email"},
		},
		verifier: provider.Verifier(&oidc.Config{ClientID: f.clientID}),
		provider: provider,
	}

	// parse the template
	app.html, err = template.ParseFiles(
		fmt.Sprintf("%v/index.tmpl", f.templatesDir),
	)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *app) index(w http.ResponseWriter, r *http.Request) {
	state := a.stateGenerator()
	a.stateStorage.add(state)

	a.html.Execute(w, map[string]string{
		"AuthURL": a.oauthConfig.AuthCodeURL(state),
	})
}

func (a *app) login(w http.ResponseWriter, r *http.Request) {
	// Validate the state
	state := r.URL.Query().Get("state")
	if !a.stateStorage.contains(state) {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}
	a.stateStorage.pop(state)

	// Exchange the code
	code := r.URL.Query().Get("code")
	ctx := r.Context()
	oauth2Token, err := a.oauthConfig.Exchange(ctx, code)
	if err != nil {
		http.Error(w, fmt.Errorf("token exchange failed: %v", err).Error(), http.StatusInternalServerError)
		return
	}

	// Verify the token
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, fmt.Errorf("no id_token field in oauth2 token").Error(), http.StatusInternalServerError)
		return
	}

	_, err = a.verifier.Verify(ctx, rawIDToken)
	if err != nil {
		http.Error(w, fmt.Errorf("ID Token verification failed: %v", err).Error(), http.StatusInternalServerError)
		return
	}

	// Get the user info from the server
	userInfo, err := a.provider.UserInfo(r.Context(), oauth2.StaticTokenSource(oauth2Token))
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to get userinfo: %v", err).Error(), http.StatusInternalServerError)
		return
	}

	userInfoJSON, err := json.MarshalIndent(userInfo, "", "	")
	if err != nil {
		http.Error(w, fmt.Errorf("Failed to get userinfo: %v", err).Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userInfoJSON)
}
