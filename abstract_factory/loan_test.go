package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewMortgageLoan(t *testing.T) {
	loan, err := NewMortgageLoan(100000.0, 10000.0, 420)
	assert.Nil(t, err)
	assert.NotNil(t, loan.LoanId)
	assert.Equal(t, loan.Amount, 100000.0)
	assert.Equal(t, loan.Income, 10000.0)
	assert.Equal(t, loan.Installments, 420)
}

func TestNewMortgageLoan_ReturnsError_WhenInstallmentsExceed420Months(t *testing.T) {
	_, err := NewMortgageLoan(100000.0, 10000.0, 520)
	assert.EqualError(t, err, "the maximum number of installments for mortgage loan is 420")
}

func TestNewMortgageLoan_InstallmentOverIncomeThreshold(t *testing.T) {
	_, err := NewMortgageLoan(200000.0, 1000.0, 420)
	assert.EqualError(t, err, "the amount installment not could exceed 25% of monthly income")
}

func TestNewCarLoan_ReturnsError_WhenInstallmentsExceed60Months(t *testing.T) {
	_, err := NewCarLoan(100000.0, 10000.0, 72)
	assert.EqualError(t, err, "the maximum number of installments for car loan is 60")
}

func TestNewCarLoan_InstallmentOverIncomeThreshold(t *testing.T) {
	_, err := NewCarLoan(200000.0, 1000.0, 60)
	assert.EqualError(t, err, "the amount installment not could exceed 30% of monthly income")
}
