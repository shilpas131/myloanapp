package loanapp

import (
	"time"
)
/*
 define the json struct body
 */
type Loan struct {
	ID      int
	Amount    string  `json:"amount"`
	StartDate time.Time `json:"start-date"`
	InterestRate string  `json:"interest-rate"`
	PaymentTracker []Payment
}

type LoanService interface {
	InitiateLoan(amount string, startDate time.Time, interestRate string) (*Loan, error)
	AddPayment(ID int, amount string, paymentDate time.Time) (bool, error)
	GetBalance(ID int, balanceDate time.Time) string
}

