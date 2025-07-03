package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMortgageLoanRequest(t *testing.T) {
	repoFactory := NewRepositoryMemoryFactory()
	loanFactory := NewMortgageFactory()
	applyForLoan := NewApplyForLoan(repoFactory, loanFactory)
	input := Input{
		Amount:      100000,
		Income:      10000,
		Installment: 240,
	}
	outputApplyForLoan, _ := applyForLoan.Execute(input)
	getLoan := NewGetLoan(repoFactory)
	outputGetLoan, _ := getLoan.Execute(InputLoan(outputApplyForLoan))
	assert.Len(t, outputGetLoan.Installments, 240)
	assert.Equal(t, 1, outputGetLoan.Installments[0].Number)
	assert.Equal(t, 1250.0, roundToTwoDecimalPlaces(outputGetLoan.Installments[0].Amount))
	assert.Equal(t, 416.67, roundToTwoDecimalPlaces(outputGetLoan.Installments[0].Amortization))
	assert.Equal(t, 833.33, roundToTwoDecimalPlaces(outputGetLoan.Installments[0].Interest))
	assert.Equal(t, 99583.33, roundToTwoDecimalPlaces(outputGetLoan.Installments[0].Balance))
	assert.Equal(t, 240, outputGetLoan.Installments[239].Number)
	assert.Equal(t, 420.14, roundToTwoDecimalPlaces(outputGetLoan.Installments[239].Amount))
	assert.Equal(t, 416.67, roundToTwoDecimalPlaces(outputGetLoan.Installments[239].Amortization))
	assert.Equal(t, 3.47, roundToTwoDecimalPlaces(outputGetLoan.Installments[239].Interest))
	assert.Equal(t, 0.0, roundToTwoDecimalPlaces(outputGetLoan.Installments[239].Balance))
}
