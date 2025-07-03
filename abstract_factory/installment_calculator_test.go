package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateInstallments_UsingSACMethod(t *testing.T) {
	installmentCalculator := SACInstallmentCalculator{}
	loan, _ := NewMortgageLoan(100000, 10000, 240)
	installments := installmentCalculator.calculate(loan)
	assert.Len(t, installments, 240)
	assert.Equal(t, 1, installments[0].Number)
	assert.Equal(t, 1250.0, roundToTwoDecimalPlaces(installments[0].Amount))
	assert.Equal(t, 416.67, roundToTwoDecimalPlaces(installments[0].Amortization))
	assert.Equal(t, 833.33, roundToTwoDecimalPlaces(installments[0].Interest))
	assert.Equal(t, 99583.33, roundToTwoDecimalPlaces(installments[0].Balance))
	assert.Equal(t, 240, installments[239].Number)
	assert.Equal(t, 420.14, roundToTwoDecimalPlaces(installments[239].Amount))
	assert.Equal(t, 416.67, roundToTwoDecimalPlaces(installments[239].Amortization))
	assert.Equal(t, 3.47, roundToTwoDecimalPlaces(installments[239].Interest))
	assert.Equal(t, 0.0, roundToTwoDecimalPlaces(installments[239].Balance))
}

func TestCalculateInstallments_UsingPriceMethod(t *testing.T) {
	installmentCalculator := PriceInstallmentCalculator{}
	loan, _ := NewMortgageLoan(100000, 10000, 240)
	installments := installmentCalculator.calculate(loan)
	assert.Len(t, installments, 240)
	assert.Equal(t, 1, installments[0].Number)
	assert.Equal(t, 965.02, roundToTwoDecimalPlaces(installments[0].Amount))
	assert.Equal(t, 131.69, roundToTwoDecimalPlaces(installments[0].Amortization))
	assert.Equal(t, 833.33, roundToTwoDecimalPlaces(installments[0].Interest))
	assert.Equal(t, 99868.31, roundToTwoDecimalPlaces(installments[0].Balance))
	assert.Equal(t, 240, installments[239].Number)
	assert.Equal(t, 965.02, roundToTwoDecimalPlaces(installments[239].Amount))
	assert.Equal(t, 957.05, roundToTwoDecimalPlaces(installments[239].Amortization))
	assert.Equal(t, 7.98, roundToTwoDecimalPlaces(installments[239].Interest))
	assert.Equal(t, 0.0, roundToTwoDecimalPlaces(installments[239].Balance))
}

func roundToTwoDecimalPlaces(v float64) float64 {
	return math.Round(v*100) / 100
}
