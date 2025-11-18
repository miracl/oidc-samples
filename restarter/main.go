package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	var host, port string
	flag.StringVar(&host, "host", "0.0.0.0", "Listen host")
	flag.StringVar(&port, "port", "8081", "Listen port")
	flag.Parse()

	http.HandleFunc("/restart", restartSignature)
	if err := http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), nil); err != nil {
		log.Fatal(err)
	}
}

func restartSignature(w http.ResponseWriter, r *http.Request) {
	containerName := r.URL.Query().Get("name")
	if containerName == "" {
		http.Error(w, "missing required query param: 'name'", http.StatusPreconditionRequired)
		return
	}

	cmd := exec.Command("docker", "restart", containerName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	log.Printf("[EXECUTING] %s\n", cmd.String())

	err := cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("docker restart error: %#v", err.Error()), http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Successfully restarted container with name: %#v", containerName)))
	w.WriteHeader(http.StatusOK)
}
