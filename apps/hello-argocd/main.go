package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	hostname string
	startAt  = time.Now()
)

func main() {
	hostname, _ = os.Hostname()
	if hostname == "" {
		hostname = "unknown"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/healthz", handleHealth)
	mux.HandleFunc("/readyz", handleReady)

	addr := ":8080"
	log.Printf("Starting hello-argocd on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"app":      "hello-argocd",
		"hostname": hostname,
		"uptime":   time.Since(startAt).Truncate(time.Second).String(),
	})
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}
