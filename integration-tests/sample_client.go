package main

import (
	"encoding/json"
	"net/http"
)

type sampleClient struct {
	authCookies []*http.Cookie
	url         string
	httpClient  *http.Client
	sessionURL  string
}

func newSampleClient(url string, httpClient *http.Client) *sampleClient {
	return &sampleClient{
		url:        url,
		httpClient: httpClient,
	}
}

func (s *sampleClient) authorize() (responseBody []byte, err error) {
	req, err := getRequest(s.url, "GET", nil)
	if err != nil {
		return nil, err
	}

	var res []byte
	res, s.authCookies, err = getResponse(req, s.httpClient)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *sampleClient) login(redirectURL string) error {
	req, err := getRequest(redirectURL, "GET", nil)
	if err != nil {
		return err
	}
	for _, cookie := range s.authCookies {
		req.AddCookie(cookie)
	}

	var sessionURL []byte
	sessionURL, s.authCookies, err = getResponse(req, s.httpClient)
	if err != nil {
		return err
	}

	s.sessionURL = string(sessionURL)
	return err
}

func (s *sampleClient) getUserInfo() (userInfo *userInfo, err error) {
	req, err := getRequest(s.sessionURL, "GET", nil)
	if err != nil {
		return nil, err
	}
	for _, cookie := range s.authCookies {
		req.AddCookie(cookie)
	}

	userInfoResponse, _, err := getResponse(req, s.httpClient)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(userInfoResponse, &userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}
