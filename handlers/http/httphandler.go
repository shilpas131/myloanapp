package http

import (
	"encoding/json"
	"log"
	"myloanapp/handlers/loan"
	loanmanager "myloanapp/handlers/loan"
	"myloanapp/loanapp"
	"net/http"
)

type Handler struct {
	LoanManager loanmanager.LoanHandler
}

func (h *Handler) InitiateLoan(w http.ResponseWriter, req *http.Request) {
	var l loanapp.Loan
	err := json.NewDecoder(req.Body).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	h.LoanManager.InitiateLoan(l)
	log.Println("New Loan Initiated ... ")
	json.NewEncoder(w).Encode(h.LoanManager.Loan)
}

func (h *Handler) AddPayment(w http.ResponseWriter, req *http.Request) {
	var p loanapp.Payment

	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	log.Println("Adding Payment ... ")
	payments := h.LoanManager.AddPayment(p)

	json.NewEncoder(w).Encode(payments)
}

func (h *Handler) GetBalance(w http.ResponseWriter, req *http.Request) {
	balanceDate := req.URL.Query().Get("balanceDate")
	log.Println("Executing GetBalance with date:" + balanceDate)
	balance, err := h.LoanManager.GetBalance(balanceDate)
	if err != nil {
		log.Println("error:Invalid request")

		http.Error(w, "Invalid date", 400) // bad request
	} else {
		log.Println("the balance is " + balance)

		w.Write([]byte("The balance is: " + balance))
	}
}

func InitNewHandlerWith(storageType string) *Handler {
	loanService := loan.GetNewLoanServiceWith(storageType)
	return &Handler{
		LoanManager: loanService,
	}
}
