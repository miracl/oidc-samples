package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

type modifier struct {
	proxyHost  string
	proxyPort  string
	httpClient *http.Client
	listener   net.Listener
	handler    func(w http.ResponseWriter, r *http.Request)
}

func newModifier(handler func(w http.ResponseWriter, r *http.Request), proxyHost, proxyPort string, httpClient *http.Client) *modifier {
	return &modifier{
		proxyHost:  proxyHost,
		proxyPort:  proxyPort,
		httpClient: httpClient,
		handler:    handler,
	}
}

func (m *modifier) start() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
		return
	}

	m.listener = l

	http.HandleFunc("/", m.handler)
	go func() {
		_ = http.Serve(m.listener, nil)
	}()

	m.setupSession()
}

func (m *modifier) getPort() int {
	return m.listener.Addr().(*net.TCPAddr).Port
}

func (m *modifier) stop() {
	if err := m.listener.Close(); err != nil {
		log.Fatal(err)
	}

	m.stopSession()
}

func (m *modifier) setupSession() {
	payload := struct {
		ModifyURL string `json:"modifyUrl"`
	}{
		ModifyURL: "http://127.0.0.1:" + fmt.Sprint(m.getPort()),
	}

	prreq, err := newRequest(fmt.Sprintf("http://%v:%v/session", m.proxyHost, m.proxyPort), "POST", payload)
	if err != nil {
		log.Fatalf("Error making request to the proxy: %v", err)
	}

	_, _, err = getResponse(prreq, m.httpClient)
	if err != nil {
		log.Fatalf("Error when receiving the proxy response: %v", err)
	}
}

func (m *modifier) stopSession() {
	prreq, err := newRequest(fmt.Sprintf("http://%v:%v/session", m.proxyHost, m.proxyPort), "DELETE", http.NoBody)
	if err != nil {
		log.Fatalf("Error when making stop session request to the proxy: %v", err)
	}

	_, _, err = getResponse(prreq, m.httpClient)
	if err != nil {
		log.Fatalf("Error on proxy response: %v", err)
	}
}
