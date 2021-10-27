package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	startServer()
}

func startServer() {
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Serving app on port %s", portNumber)
}
