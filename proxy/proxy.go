package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/elazarl/goproxy"
)

type proxy struct {
	sessionMux sync.RWMutex
	session    *session
}

type session struct {
	MediatorURL string `json:"mediatorUrl"`
}

func (p *proxy) setupProxy() *goproxy.ProxyHttpServer {
	proxy := goproxy.NewProxyHttpServer()

	proxy.NonproxyHandler = p.sessionHandler()

	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnResponse().DoFunc(p.modifyResponse)

	return proxy
}

func (p *proxy) sessionHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/session" {
			return
		}
		if req.Method == "POST" && p.session == nil {
			p.startSession(req)
		}
		if req.Method == "DELETE" {
			p.sessionMux.Lock()
			p.session = nil
			p.sessionMux.Unlock()
		}
	})
}

func (p *proxy) startSession(req *http.Request) error {
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	defer req.Body.Close()

	var sessionResp session
	if err = json.Unmarshal(reqBody, &sessionResp); err != nil {
		return err
	}

	p.sessionMux.Lock()
	p.session = &session{
		MediatorURL: sessionResp.MediatorURL,
	}
	p.sessionMux.Unlock()

	return nil
}

func (p *proxy) modifyResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	if p.session != nil {
		p.sessionMux.RLock()
		modifiedResponse, err := http.Post(p.session.MediatorURL, "application/json", resp.Body)
		p.sessionMux.RUnlock()
		if err != nil {
			log.Fatal("Error POSTing to mediator", err)
		}

		resp.Body = ioutil.NopCloser(modifiedResponse.Body)
	}

	return resp
}
