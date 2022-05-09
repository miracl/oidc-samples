package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/modifySignature", modifySignature)

	var host, port string
	flag.StringVar(&host, "host", "0.0.0.0", "Listen host")
	flag.StringVar(&port, "port", "8081", "Listen port")
	flag.Parse()

	address := host + ":" + port

	log.Fatal(http.ListenAndServe(address, nil))
}

func modifySignature(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.Write([]byte(fmt.Sprintf("ERROR is: %#v", err.Error())))
		return
	}

	originalRequestURL := r.Header.Get("X-Forwarded-Host")
	if originalRequestURL == "https://api.mpin.io:443/oidc/certs" {
		keys := struct {
			Keys []map[string]string `json:"keys"`
		}{}

		if err = json.Unmarshal(body, &keys); err != nil {
			w.Write([]byte(fmt.Sprintf("/oidc/certs keys decoding error: %#v", err.Error())))
			return
		}

		keys.Keys[0]["n"] = "invalid"
		keys.Keys[1]["n"] = "invalid"

		body, err = json.Marshal(keys)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("keys encoding error: %#v", err.Error())))
			return
		}
	}

	for k, v := range r.Header {
		if k != "X-Forwarded-Host" {
			for _, hv := range v {
				w.Header().Set(k, hv)
			}
		}
	}

	w.Write(body)
}
