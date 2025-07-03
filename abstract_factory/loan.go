package main

import (
	"crypto/rand"
	"errors"

	"github.com/google/uuid"
)

type Loan interface {
	GetLoanId() string
	GetAmount() float64
	GetIncome() float64
	GetInstallments() int
	GetRate() float64
}

type LoanBase struct {
	LoanId       string
	Amount       float64
	Income       float64
	Installments int
	Rate         float64
}

func (lb LoanBase) GetLoanId() string {
	return lb.LoanId
}
func (lb LoanBase) GetAmount() float64 {
	return lb.Amount
}
func (lb LoanBase) GetIncome() float64 {
	return lb.Income
}
func (lb LoanBase) GetInstallments() int {
	return lb.Installments
}
func (lb LoanBase) GetRate() float64 {
	return lb.Rate
}

type MortgageLoan struct {
	LoanBase
}

func NewMortgageLoan(amount, income float64, installments int) (*MortgageLoan, error) {
	loanId := uuid.New().String()
	if installments > 420 {
		return &MortgageLoan{}, errors.New("the maximum number of installments for mortgage loan is 420")
	}
	if (income * 0.25) < amount/float64(installments) {
		return &MortgageLoan{}, errors.New("the amount installment not could exceed 25% of monthly income")
	}
	return &MortgageLoan{LoanBase{loanId, amount, income, installments, 10}}, nil
}

type CarLoan struct {
	LoanBase
}

func NewCarLoan(amount, income float64, installments int) (*CarLoan, error) {
	loanId := rand.Text()
	if installments > 60 {
		return &CarLoan{}, errors.New("the maximum number of installments for car loan is 60")
	}
	if (income * 0.30) < amount/float64(installments) {
		return &CarLoan{}, errors.New("the amount installment not could exceed 30% of monthly income")
	}
	return &CarLoan{LoanBase{loanId, amount, income, installments, 15}}, nil
}
