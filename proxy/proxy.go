package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/elazarl/goproxy"
)

func setupProxy() *goproxy.ProxyHttpServer {
	proxy := goproxy.NewProxyHttpServer()

	proxy.NonproxyHandler = sessionHandler()

	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnResponse().DoFunc(modifyResponse)

	return proxy
}

func modifyResponse(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	if currentSession.IsActive {
		modifiedResponse, err := post(currentSession.MediatorURL, resp.Body)
		if err != nil {
			log.Fatal("Error POSTing to mediator", err)
		}

		resp.Body = ioutil.NopCloser(bytes.NewBuffer(modifiedResponse))
	}

	return resp
}
