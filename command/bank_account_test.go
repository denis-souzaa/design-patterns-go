package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustMakeTransaction(t *testing.T) {
	t.Run("Deve fazer uma transferencia bancária", func(t *testing.T) {
		from := &BankAccount{}
		to := &BankAccount{}
		assert.Equal(t, float64(0), from.Balance())
		assert.Equal(t, float64(0), to.Balance())
		from.Debit(100)
		to.Credit(100)
		assert.Equal(t, float64(-100), from.Balance())
		assert.Equal(t, float64(100), to.Balance())
	})

	t.Run("Deve fazer um transferencia bancária usando command", func(t *testing.T) {
		from := &BankAccount{}
		to := &BankAccount{}
		assert.Equal(t, float64(0), from.Balance())
		assert.Equal(t, float64(0), to.Balance())
		transferCommand := NewTransferCommand(from, to, 100)
		transferCommand.Execute()
		assert.Equal(t, float64(-100), from.Balance())
		assert.Equal(t, float64(100), to.Balance())
	})
}
