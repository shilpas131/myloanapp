package main

import (
	loanhttphandler "myloanapp/handlers/http"
	"net/http"
)

func main() {

	// creating new http handler with in-memory storage.
	var h = loanhttphandler.InitNewHandlerWith("inMemoryStorage")

	http.HandleFunc("/initiateLoan", h.InitiateLoan)
	http.HandleFunc("/addPayment", h.AddPayment)
	http.HandleFunc("/getBalance", h.GetBalance)

	loanhttphandler.ConfigureAndStartServer()
}