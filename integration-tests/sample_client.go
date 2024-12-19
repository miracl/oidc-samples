package main

import (
	"fmt"
	"net"
	"net/http"
)

type sampleClient struct {
	cookies    []*http.Cookie
	url        string
	httpClient *http.Client
}

func newSampleClient(url string, httpClient *http.Client) *sampleClient {
	return &sampleClient{
		url:        url,
		httpClient: httpClient,
	}
}

func (s *sampleClient) ping() error {
	req, err := http.NewRequest("HEAD", s.url, http.NoBody)
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

func (s *sampleClient) restart(restarterHost, restarterPort, sampleName string) error {
	restartSampleReq, err := newRequest(
		fmt.Sprintf("http://%s/restart?name=%s", net.JoinHostPort(restarterHost, restarterPort), sampleName),
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
