package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustMakeTransferBank(t *testing.T) {
	bankAccountRepo := &BankAccountRepositoryMemory{}
	bankAccountRepo.Save(BankAccount{BankAccountId: 1})
	bankAccountRepo.Save(BankAccount{BankAccountId: 2})
	makeTransfer := NewMakeTransfer(bankAccountRepo)
	input := Input{
		FromBankAccountId: 1,
		ToBankAccountId:   2,
		Amount:            100,
	}
	makeTransfer.Execute(input)
	balance := NewBalance(bankAccountRepo)
	outputA := balance.Execute(1)
	assert.Equal(t, float64(-100), outputA.Balance)
	outputB := balance.Execute(2)
	assert.Equal(t, float64(100), outputB.Balance)
}
