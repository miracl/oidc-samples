package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var errInvalidLogin = errors.New("invalid login")

type sampleClient struct {
	cookies    []*http.Cookie
	url        string
	httpClient *http.Client
	sessionURL string
}

func newSampleClient(url string, httpClient *http.Client) *sampleClient {
	return &sampleClient{
		url:        url,
		httpClient: httpClient,
	}
}

func (s *sampleClient) ping() error {
	req, err := http.NewRequest("HEAD", s.url, nil)
	if err != nil {
		return err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()

	return nil
}

func (s *sampleClient) authorize() (responseBody []byte, err error) {
	req, err := newRequest(s.url, "GET", nil)
	if err != nil {
		return nil, err
	}

	var res []byte
	res, s.cookies, err = getResponse(req, s.httpClient)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *sampleClient) login(redirectURL string) error {
	req, err := newRequest(redirectURL, "GET", nil)
	if err != nil {
		return err
	}

	s.addCookies(req)

	var sessionURL []byte
	sessionURL, s.cookies, err = getResponse(req, s.httpClient)
	if err != nil {
		return err
	}

	_, err = url.Parse(string(sessionURL))
	if err != nil {
		return fmt.Errorf("%w: %s", errInvalidLogin, err.Error())
	}

	s.sessionURL = string(sessionURL)
	return err
}

func (s *sampleClient) getUserInfo() (userInfo *userInfo, err error) {
	req, err := newRequest(s.sessionURL, "GET", nil)
	if err != nil {
		return nil, err
	}

	s.addCookies(req)
	userInfoResponse, _, err := getResponse(req, s.httpClient)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(userInfoResponse, &userInfo); err != nil {
		return nil, err
	}

	return userInfo, nil
}

func (s *sampleClient) addCookies(req *http.Request) {
	for _, cookie := range s.cookies {
		req.AddCookie(cookie)
	}
}

func (s *sampleClient) restart(restarterHost, restarterPort, sampleName string) error {
	restartSampleReq, err := newRequest(
		fmt.Sprintf("http://%s:%s/restart?name=%s", restarterHost, restarterPort, sampleName),
		"POST",
		http.NoBody)
	if err != nil {
		return err
	}

	_, _, err = getResponse(restartSampleReq, s.httpClient)
	if err != nil {
		return err
	}

	for {
		if err = s.ping(); err == nil {
			break
		}
	}

	return nil
}
