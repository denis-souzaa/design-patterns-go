package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustMakeTransaction(t *testing.T) {
	from := &BankAccount{}
	to := &BankAccount{}
	assert.Equal(t, float64(0), from.Balance())
	assert.Equal(t, float64(0), to.Balance())
	from.Debit(100)
	to.Credit(100)
	assert.Equal(t, float64(-100), from.Balance())
	assert.Equal(t, float64(100), to.Balance())
}

func TestMustMakeTransactionWithCommand(t *testing.T) {
	from := &BankAccount{}
	to := &BankAccount{}
	assert.Equal(t, float64(0), from.Balance())
	assert.Equal(t, float64(0), to.Balance())
	transferCommand := NewTransferCommand(from, to, 100)
	transferCommand.Execute()
	assert.Equal(t, float64(-100), from.Balance())
	assert.Equal(t, float64(100), to.Balance())
}
