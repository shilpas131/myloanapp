package loanapp

import "time"

type Payment struct {
	PaymentAmount    string `json:"amount"`
	PaymentDate time.Time `json:"date"`
}

type ByDate []Payment

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Less(i, j int) bool { return a[i].PaymentDate.Before(a[j].PaymentDate) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
