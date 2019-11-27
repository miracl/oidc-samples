package main

import (
	"flag"
	"log"
	"net/http"
)

type flags struct {
	addr         string
	templatesDir string
	clientID     string
	clientSecret string
	redirectURL  string
	issuer       string
}

func main() {
	flags := parseFlags()

	app, err := newApp(flags)
	if err != nil {
		log.Fatal(err)
	}

	startServer(app, flags)
}

func parseFlags() *flags {
	f := &flags{}

	flag.StringVar(&f.addr, "addr", ":8000", "Listen address")
	flag.StringVar(&f.templatesDir, "templates-dir", "templates", "Template files location")
	flag.StringVar(&f.clientID, "client-id", "", "OIDC Client Id")
	flag.StringVar(&f.clientSecret, "client-secret", "", "OIDC Client Secret")
	flag.StringVar(&f.redirectURL, "redirect-url", "http://localhost:8000/login", "Redirect URL")
	flag.StringVar(&f.issuer, "issuer", "https://api.mpin.io", "Backend url")
	flag.Parse()

	return f
}

func startServer(app *app, config *flags) {
	http.HandleFunc("/", app.index)
	http.HandleFunc("/login", app.login)
	log.Printf("Server started. Listening on %v", config.addr)
	log.Fatal(http.ListenAndServe(config.addr, nil))
}
