package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	setCA(caCert, caKey)
	proxy := setupProxy()

	var host, port string
	flag.StringVar(&host, "host", "0.0.0.0", "Listen host")
	flag.StringVar(&port, "port", "8080", "Listen port")
	verbose := flag.Bool("v", false, "Log every proxied request to stdout")
	flag.Parse()

	proxy.Verbose = *verbose
	address := host + ":" + port

	log.Fatal(http.ListenAndServe(address, proxy))

}
