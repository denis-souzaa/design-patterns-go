package command

type TransferCommand struct {
	From   *BankAccount
	To     *BankAccount
	Amount float64
}

func NewTransferCommand(from, to *BankAccount, amount float64) *TransferCommand {
	return &TransferCommand{From: from, To: to, Amount: amount}
}

func (tc *TransferCommand) Execute() {
	tc.From.Debit(tc.Amount)
	tc.To.Credit(tc.Amount)
}
