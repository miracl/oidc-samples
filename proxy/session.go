package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type session struct {
	IsActive    bool
	MediatorURL string `json:"mediatorUrl"`
}

var currentSession session

func sessionHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path == "/session" {
			if req.Method == "POST" {
				startSession(req)
			}
			if req.Method == "DELETE" {
				currentSession.IsActive = false
			}
		}
	})
}

func startSession(req *http.Request) error {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	var sessionResp session
	if err = json.Unmarshal(reqBody, &sessionResp); err != nil {
		return err
	}

	currentSession = session{
		IsActive:    true,
		MediatorURL: sessionResp.MediatorURL,
	}

	return nil
}
