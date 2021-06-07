package http

import (
	"log"
	"net/http"
)

func ConfigureAndStartServer() {
	server := http.Server{
		Addr:    "localhost:4200",
	}

	log.Println("Starting Server ....")
	server.ListenAndServe()
}
