package command

type Output struct {
	Balance float64
}
type Balance struct {
	BankAccountRepo BankAccountRepository
}

func NewBalance(bankAccountRepo BankAccountRepository) *Balance {
	return &Balance{BankAccountRepo: bankAccountRepo}
}

func (mt *Balance) Execute(bankAccountId int) Output {
	bankAccount := mt.BankAccountRepo.ById(bankAccountId)
	return Output{Balance: bankAccount.Balance()}
}
