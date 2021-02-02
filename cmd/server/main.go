package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	listenAddress := ":8080"
	log.Printf("Starting server @ %s", listenAddress)
	http.HandleFunc("/", logger(handler))
	log.Panic(http.ListenAndServe(listenAddress, nil))
}

func logger(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.String())
		f(w,r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "Hello from "+r.URL.String())
	if err != nil {
		log.Printf("Error %v", err)
	}
}
