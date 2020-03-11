package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func post(url string, payload io.Reader) (responseBody []byte, err error) {
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	c := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
