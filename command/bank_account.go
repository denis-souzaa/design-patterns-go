package command

type BankAccount struct {
	BankAccountId int
	Amount        float64
	Transactions  []Transaction
}

func (ba *BankAccount) Debit(amount float64) {
	ba.Transactions = append(ba.Transactions, Transaction{Type: Debit, Amount: amount})
}

func (ba *BankAccount) Credit(amount float64) {
	ba.Transactions = append(ba.Transactions, Transaction{Type: Credit, Amount: amount})
}

func (ba *BankAccount) Balance() float64 {
	var total float64
	for _, t := range ba.Transactions {
		if t.Type == Credit {
			total += t.Amount
		}

		if t.Type == Debit {
			total -= t.Amount
		}
	}

	return total
}
