package main

import "math"

type InstallmentCalculator interface {
	calculate(loan Loan) []Installment
}

type SACInstallmentCalculator struct {
}

func NewSACInstallmentCalculator() *SACInstallmentCalculator {
	return &SACInstallmentCalculator{}
}

func (sic *SACInstallmentCalculator) calculate(loan Loan) []Installment {
	installments := []Installment{}
	balance := loan.GetAmount()
	rate := (loan.GetRate() / 100) / 12
	installmentNumber := 1
	amortization := balance / float64(loan.GetInstallments())
	for balance > 0.10 {
		interest := balance * rate
		updatedBalance := balance + interest
		amount := interest + amortization
		balance = updatedBalance - amount
		if balance < 0.10 {
			balance = 0
		}
		installment := Installment{loan.GetLoanId(), installmentNumber, amount, amortization, interest, balance}
		installments = append(installments, installment)
		installmentNumber++
	}
	return installments
}

type PriceInstallmentCalculator struct {
}

func NewPriceInstallmentCalculator() *PriceInstallmentCalculator {
	return &PriceInstallmentCalculator{}
}

func (pc *PriceInstallmentCalculator) calculate(loan Loan) []Installment {
	installments := []Installment{}
	balance := loan.GetAmount()
	rate := (loan.GetRate() / 100) / 12
	installmentNumber := 1
	formula := math.Pow((1 + rate), float64(loan.GetInstallments()))
	amount := balance * ((formula * rate) / (formula - 1))
	for balance > 2 {
		interest := balance * rate
		amortization := amount - interest
		balance -= amortization
		if balance < 2 {
			balance = 0
		}
		installment := Installment{loan.GetLoanId(), installmentNumber, amount, amortization, interest, balance}
		installments = append(installments, installment)
		installmentNumber++
	}
	return installments
}
