package main

import (
	"assignments/4/router"
	"log"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.GetRouter(),
	}
	log.Println("Listening...")
	server.ListenAndServe()
}
