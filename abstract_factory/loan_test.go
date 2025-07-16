package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewMortgageLoan(t *testing.T) {
	t.Run("Deve criar um novo financiamento imobiliario", func(t *testing.T) {
		loan, err := NewMortgageLoan(100000.0, 10000.0, 420)
		assert.Nil(t, err)
		assert.NotNil(t, loan.LoanId)
		assert.Equal(t, loan.Amount, 100000.0)
		assert.Equal(t, loan.Income, 10000.0)
		assert.Equal(t, loan.Installments, 420)
	})

	t.Run("Não deve criar um novo funciamento se o numero de parcelas for maior que 420", func(t *testing.T) {
		_, err := NewMortgageLoan(100000.0, 10000.0, 520)
		assert.EqualError(t, err, "the maximum number of installments for mortgage loan is 420")
	})

	t.Run("Não deve criar um novo financiamento se o valor da parcela for maior que 25% da renda", func(t *testing.T) {
		_, err := NewMortgageLoan(200000.0, 1000.0, 420)
		assert.EqualError(t, err, "the amount installment not could exceed 25% of monthly income")
	})
}

func TestNewCarLoan(t *testing.T) {
	t.Run("Não deve criar financiamento veicular se número de parcela for maior que 60", func(t *testing.T) {
		_, err := NewCarLoan(100000.0, 10000.0, 72)
		assert.EqualError(t, err, "the maximum number of installments for car loan is 60")
	})

	t.Run("Não deve criar financiamento veicular se valor da parcela for maior que 30% da renda", func(t *testing.T) {
		_, err := NewCarLoan(200000.0, 1000.0, 60)
		assert.EqualError(t, err, "the amount installment not could exceed 30% of monthly income")
	})
}
