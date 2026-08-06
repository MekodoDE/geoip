package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type CheckRequest struct {
	Attributes struct {
		Request struct {
			HTTP struct {
				Host   string `json:"host"`
				Method string `json:"method"`
				Path   string `json:"path"`
			} `json:"http"`
		} `json:"request"`
	} `json:"attributes"`
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	var req CheckRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf(
		"authorization request host=%s method=%s path=%s",
		req.Attributes.Request.HTTP.Host,
		req.Attributes.Request.HTTP.Method,
		req.Attributes.Request.HTTP.Path,
	)

	// Aktuell: alles erlauben
	// Später kommt hier GeoIP-Prüfung rein.

	w.WriteHeader(http.StatusOK)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/check", checkHandler)
	http.HandleFunc("/healthz", healthHandler)

	log.Println("geoip-auth listening on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
