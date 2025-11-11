package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var options struct {
	clientID        string
	clientSecret    string
	redirectURL     string
	projectDomain   string
	sampleURL       string
	proxyHost       string
	proxyPort       string
	restarterHost   string
	restarterPort   string
	sampleName      string
	skipModifyTests bool
}

func TestMain(m *testing.M) {
	flag.StringVar(&options.clientID, "client-id", "", "the client id of your MIRACL Trust Portal application")
	flag.StringVar(&options.clientSecret, "client-secret", "", "the client secret of your MIRACL Trust Portal application")
	flag.StringVar(&options.redirectURL, "redirect-url", "http://localhost:8000/login", "the redirect url of your MIRACL Trust Portal application")
	flag.StringVar(&options.projectDomain, "project-domain", "", "the project domain of your MIRACL Trust Portal application")
	flag.StringVar(&options.sampleURL, "sample-url", "http://127.0.0.1:8000", "the sample URL")
	flag.StringVar(&options.proxyHost, "proxy-host", "", "Sample's proxy HOST")
	flag.StringVar(&options.proxyPort, "proxy-port", "", "Sample's proxy PORT")
	flag.StringVar(&options.restarterHost, "restarter-host", "127.0.0.1", "Restarter's HOST")
	flag.StringVar(&options.restarterPort, "restarter-port", "8081", "Restarter's PORT")
	flag.StringVar(&options.sampleName, "sample-name", "sample", "Sample's container name")
	flag.BoolVar(&options.skipModifyTests, "skip-modify-tests", true, "Specify if the tests which modify requests to be skipped")

	flag.Parse()

	if options.projectDomain == "" {
		fmt.Println("ERROR: project-domain is missing.\nUse -h flag to see all args.")
	}

	if options.clientSecret == "" && options.clientID == "" {
		fmt.Println("ERROR: client-id and client-secret args are missing.\nUse -h flag to see all args.")
		os.Exit(1)
	}

	os.Exit(m.Run())
}
