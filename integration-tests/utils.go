package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	mathRand "math/rand"
	"net/http"
	"strconv"
	"time"

	amclRand "code.miracl.com/maas/maas/src/lib/gomiracl/rand"
)

func newRequest(url, method string, payload interface{}, headers ...header) (req *http.Request, err error) {
	if method == "GET" {
		req, err = http.NewRequest(method, url, http.NoBody)
	} else {
		reqPayloadJSON, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method, url, bytes.NewReader(reqPayloadJSON))
		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	for _, h := range headers {
		req.Header.Add(h.Key, h.Value)
	}

	return req, err
}

func getResponse(req *http.Request, httpClient *http.Client) (responseBody []byte, cookies []*http.Cookie, err error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	// Whenever we're redirected we take the Location and return it.
	if resp.StatusCode == http.StatusFound || resp.StatusCode == http.StatusMovedPermanently {
		redirectLocation, err := resp.Location()
		if err != nil {
			return nil, nil, err
		}

		return []byte(redirectLocation.String()), resp.Cookies(), nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, nil, fmt.Errorf("unsuccessful request (%s) to %s", strconv.Itoa(resp.StatusCode), req.URL.String())
	}

	responseBody, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return responseBody, resp.Cookies(), nil
}

func makeRequest(httpClient *http.Client, url, method string, payload interface{}, headers ...header) (responseBody []byte, err error) {
	req, err := newRequest(url, method, payload, headers...)
	if err != nil {
		return nil, err
	}

	res, _, err := getResponse(req, httpClient)

	return res, err
}

func randPIN() int {
	mathRand.Seed(time.Now().UnixNano())

	return mathRand.Intn(9000) + 1000
}

func newRand() *amclRand.Rand {
	r, err := amclRand.New(rand.Reader, 128)
	if err != nil {
		panic(err)
	}

	return r
}
