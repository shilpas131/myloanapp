package loan

import (
	. "github.com/smartystreets/goconvey/convey"
	"myloanapp/loanapp"
	"testing"
	"time"
)

var testCasesForLoan = []loanapp.Loan{
	{
		ID : 23,
		Amount : "300000",
		StartDate : time.Now(),
		InterestRate : "7",
		PaymentTracker : []loanapp.Payment{},
	},{
		ID : 23,
		Amount : "400000",
		StartDate : time.Now(),
		InterestRate : "10",
		PaymentTracker : []loanapp.Payment{},
	},{
		ID : 23,
		Amount : "30000",
		StartDate : time.Now(),
		InterestRate : "4",
		PaymentTracker : []loanapp.Payment{},
	},
}

var testCasesForPayment = []loanapp.Payment{
	{
		PaymentAmount : "100000",
		PaymentDate : time.Now(),
	},{
		PaymentAmount : "40000",
		PaymentDate : time.Now().Add(24),
	},{
		PaymentAmount : "5000",
		PaymentDate : time.Now().Add(168),
	},
}

func TestInitiateLoan(t *testing.T) {
	lh := GetNewLoanServiceWith("inMemoryStorage")

	for _, tc := range testCasesForLoan {
		Convey("When given a handler, initiate multiple loans", t, func() {
			lh.InitiateLoan(tc)
			So(lh.Loan.Amount, ShouldEqual, tc.Amount)
			So(lh.Loan.StartDate, ShouldEqual, tc.StartDate)
			So(lh.Loan.InterestRate, ShouldEqual, tc.InterestRate)
		})
	}
}

func TestAddPayment(t *testing.T) {
	lh := GetNewLoanServiceWith("inMemoryStorage")
	lh.Loan = loanapp.Loan{
		ID : 23,
		Amount : "300000",
		StartDate : time.Now(),
		InterestRate : "4",
		PaymentTracker : []loanapp.Payment{},
	}

	for in, tc := range testCasesForPayment {
		Convey("When given a handler and loan, add more payments", t, func() {
			lh.AddPayment(tc)
			So(lh.Loan.PaymentTracker[in].PaymentAmount, ShouldEqual, tc.PaymentAmount)
			So(lh.Loan.PaymentTracker[in].PaymentDate, ShouldEqual, tc.PaymentDate)
		})
	}
}

func TestBalance(t *testing.T) {
	lh := GetNewLoanServiceWith("inMemoryStorage")
	loan := loanapp.Loan{
		ID : 23,
		Amount : "300000",
		StartDate : time.Now(),
		InterestRate : "4",
		PaymentTracker : []loanapp.Payment{},
	}
	lh.InitiateLoan(loan)
	paymentList := []loanapp.Payment{
		{
	PaymentAmount : "100000",
		PaymentDate : time.Now(),
	},{
	PaymentAmount : "40000",
		PaymentDate : time.Now().Add(24),
	}}

	lh.AddPayment(paymentList[0])

	lh.AddPayment(paymentList[1])
		Convey("When given a handler, loan, and payments", t, func() {
			Convey("When queried date is before loan start date", func() {
				bal, err := lh.GetBalance("2015-05-17")
				So(err, ShouldNotBeNil)
				So(bal, ShouldEqual, "")
			})
			Convey("When queried date is after loan start date", func() {
				bal, err := lh.GetBalance("2022-05-17")
				So(err, ShouldBeNil)
				So(bal, ShouldEqual, "160000")
			})
		})
}