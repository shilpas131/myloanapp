package loan

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"myloanapp/handlers/storage/inmemorystorage"
	"myloanapp/loanapp"
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

type LoanHandler struct {
	Loan            loanapp.Loan
	storageInstance loanapp.StorageService
}

func (s *LoanHandler) InitiateLoan(loan loanapp.Loan) {
	s.Loan = loan
	randomID := rand.Intn(100-0) + 1
	s.Loan.ID = randomID
	s.storageInstance.AddNewEntryToStorage(randomID, loan)
	log.Printf("New Loan ID :  %+v\n", s.Loan.ID)
	log.Printf(" Loan : %+v\n", s.Loan)
}

func (s *LoanHandler) AddPayment(payment loanapp.Payment) []loanapp.Payment {
	addedPayment := append(s.Loan.PaymentTracker, payment)
	s.Loan.PaymentTracker = addedPayment
	fmt.Printf(" Added Payment : %+v\n", addedPayment)
	s.storageInstance.UpdateEntryInStorage(s.Loan.ID, s.Loan)
	log.Printf("Added new Payment for:  %+v\n", s.Loan.ID)
	fmt.Printf(" Loan : %+v\n", s.Loan)

	return addedPayment
}

func (s *LoanHandler) GetBalance(queryDate string) (string, error) {
	log.Printf("Getting balance for:  %+v\n", s.Loan.ID)

	parsedDate, err := time.Parse("2006-01-02", queryDate)
	if err != nil {
		log.Println(err.Error() + "error:Parsing Date failed")

		return "", err
	}
	fmt.Printf("parsedDate : %+v\n", parsedDate)
	loanData, err := s.storageInstance.GetEntryFromStorage(s.Loan.ID)
	if err != nil {
		log.Printf("error:no data in cache:  %+v\n", s.Loan.ID)

		return "", err
	}
	loan := loanData.(*loanapp.Loan)
	if parsedDate.Before(loan.StartDate) {
		log.Printf("error:queried date is before loan start date:  %+v\n", s.Loan.StartDate)

		return "", errors.New("queried date is before loan start date")
	}

	balance, err := calculateBalance(loan, parsedDate)

	return balance, err
}

func GetNewLoanServiceWith(storagetype string) LoanHandler {
	storage := inmemorystorage.GetNewInMemoryStorageLoanHandler()
	pi := new(loanapp.StorageService)
	*pi = storage

	return LoanHandler{
		storageInstance: storage,
	}
}

func calculateBalance(loanData *loanapp.Loan, parsedDate time.Time) (string, error) {
	paymentList := loanData.PaymentTracker
	sort.Sort(loanapp.ByDate(paymentList))
	principleAmount, err := decimal.NewFromString(loanData.Amount)
	interestRate, err := decimal.NewFromString(loanData.InterestRate)
	startDate := loanData.StartDate
	var previousPaymentDate time.Time
	if err != nil {
		log.Printf("error:parsing principle amount din work %v", loanData.ID)

		return "", err
	}
	// this is
	carryForwardInterest := getInterestFor(parsedDate, startDate, principleAmount, interestRate)
	balance := principleAmount.Add(carryForwardInterest)
	log.Printf("initial Interest and balance : %+v , %+v", carryForwardInterest, balance)

	for _, k := range paymentList {
		if k.PaymentDate == parsedDate || k.PaymentDate.Before(parsedDate) {

			log.Printf("parsedDate %v PaymentDate %v", parsedDate, k.PaymentDate)

			payedAmount, _ := decimal.NewFromString(k.PaymentAmount)

			log.Printf("previousPaymentDate %v", previousPaymentDate)

			if previousPaymentDate.IsZero() {
				interestNow := getInterestFor(k.PaymentDate, startDate, principleAmount,
					interestRate)
				balance = principleAmount.Add(interestNow).
					Sub(payedAmount)
				log.Printf("carryForwardInterest and balance : %+v , %+v", interestNow, balance)
				if interestNow.LessThan(payedAmount) {
					principleAmount = principleAmount.Sub(payedAmount.Sub(interestNow))
					carryForwardInterest = decimal.NewFromInt32(0)
				} else {
					carryForwardInterest = interestNow.Sub(payedAmount)
				}
				log.Printf("carryForwardInterest and principleAmount : %+v , %+v",
					carryForwardInterest, principleAmount)
			} else {
				interestNow := getInterestFor(k.PaymentDate, previousPaymentDate, principleAmount,
					interestRate)
				balance = principleAmount.Add(carryForwardInterest).Add(interestNow).
					Sub(payedAmount)
				if interestNow.Add(carryForwardInterest).LessThan(payedAmount) {
					principleAmount = principleAmount.Sub(payedAmount.Sub(interestNow))
					carryForwardInterest = decimal.NewFromInt32(0)
				} else {
					carryForwardInterest = interestNow.Add(carryForwardInterest).Sub(payedAmount)
				}
			}
		} else {
			log.Printf("Done with calculating balance %v", loanData.ID)

			return balance.Round(2).String(), nil
		}

		previousPaymentDate = k.PaymentDate
	}
	listLen := len(paymentList)
	// this is when the querydate is after the last payment date
	if listLen != 0 && paymentList[listLen-1].PaymentDate.Before(parsedDate) {
		p := paymentList[listLen-1]
		interestNow := getInterestFor(parsedDate, p.PaymentDate, principleAmount,
			interestRate)
		balance = principleAmount.Add(carryForwardInterest).Add(interestNow)
	}

	return balance.Round(2).String(), nil
}

func getInterestFor(calcDate time.Time, startDate time.Time,
	principleAmount decimal.Decimal, interestRate decimal.Decimal) decimal.Decimal {
	days := calcDate.Sub(startDate)
	b := decimal.NewFromFloat(days.Hours() / 24)

	return principleAmount.Mul(interestRate).Mul(b).
		Div(decimal.NewFromInt32(365)).Div(decimal.NewFromInt32(100))
}
