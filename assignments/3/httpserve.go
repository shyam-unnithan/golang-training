package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprintf(w, "GoLang http server <a href='/welcome'>welcome</a>")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "text/html")
	fmt.Fprintf(w, "Welcome to GoLang")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(index))
	mux.HandleFunc("/welcome", welcome)

	log.Println(".... Listening ....")

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
