package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	hostname string
)

func main() {
	listenAddress := ":8080"
	hostname, _ = os.Hostname()

	log.Printf("Starting server @ %s", listenAddress)
	http.HandleFunc("/", logger(handler))
	log.Panic(http.ListenAndServe(listenAddress, nil))
}

func logger(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", hostname, r.Method, r.URL.String())
		f(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprintf(w, "Hello from %s -> %s %s",
		hostname, r.Method, r.URL.String())
	if err != nil {
		log.Printf("Error %v", err)
	}
}
