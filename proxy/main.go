package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/elazarl/goproxy"
)

type session struct {
	RedirectToURL string `json:"redirectToUrl"`
	IsStarted     bool   `json:"isStarted"`
}

var currentSession session

func main() {
	setCA(caCert, caKey)

	proxy := goproxy.NewProxyHttpServer()

	proxy.NonproxyHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/session" {
			if r.Method == "POST" {
				startSession(r)
			}
			if r.Method == "DELETE" {
				stopSession()
			}
		}
	})

	proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)

	proxy.OnRequest(goproxy.ReqHostMatches(regexp.MustCompile("^.*$"))).DoFunc(func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
		return req, ctx.Resp
	})

	proxy.OnResponse().DoFunc(func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
		// if ctx.Req.URL.Path == "/oidc/token" {
		// 	var csResponse *testResponse
		// 	responseBody, _ := ioutil.ReadAll(resp.Body)
		// 	ctx.Logf("before change %v", string(responseBody))

		//  csResponse must have the exact same structure as the json you want to change
		// 	if err := json.Unmarshal(responseBody, &csResponse); err != nil {
		// 		ctx.Logf("%v", err)
		// 	}
		// 	// csResponse.AccessToken = "change here"
		// 	wat, _ := json.Marshal(csResponse)
		// 	resp.Body = ioutil.NopCloser(bytes.NewBuffer(wat))
		// }

		return resp
	})

	var host, port string
	flag.StringVar(&host, "host", "0.0.0.0", "Listen host")
	flag.StringVar(&port, "port", "8080", "Listen port")
	verbose := flag.Bool("v", false, "Log every proxied request to stdout")
	flag.Parse()

	proxy.Verbose = *verbose
	address := host + ":" + port

	log.Fatal(http.ListenAndServe(address, proxy))
}

func startSession(r *http.Request) {

}

func stopSession() {
	fmt.Println("TODO: Implement")
}
