package command

type BankAccountRepository interface {
	Save(bankAccount BankAccount)
	Update(bankAccount BankAccount)
	ById(bankAccountId int) *BankAccount
}

type BankAccountRepositoryMemory struct {
	bankAccounts []BankAccount
}

func (bam *BankAccountRepositoryMemory) Save(ba BankAccount) {
	bam.bankAccounts = append(bam.bankAccounts, ba)
}

func (bam *BankAccountRepositoryMemory) Update(ba BankAccount) {
	for i, c := range bam.bankAccounts {
		if c.BankAccountId == ba.BankAccountId {
			bam.bankAccounts[i].Amount = ba.Amount
			bam.bankAccounts[i].Transactions = ba.Transactions
			break
		}
	}
}

func (bam *BankAccountRepositoryMemory) ById(baId int) *BankAccount {
	for _, ba := range bam.bankAccounts {
		if ba.BankAccountId == baId {
			return &ba
		}
	}
	return nil
}
