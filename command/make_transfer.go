package command

type Input struct {
	FromBankAccountId int
	ToBankAccountId   int
	Amount            float64
}
type MakeTransfer struct {
	BankAccountRepo BankAccountRepository
}

func NewMakeTransfer(bankAccountRepo BankAccountRepository) *MakeTransfer {
	return &MakeTransfer{BankAccountRepo: bankAccountRepo}
}

func (mt *MakeTransfer) Execute(i Input) {
	from := mt.BankAccountRepo.ById(i.FromBankAccountId)
	to := mt.BankAccountRepo.ById(i.ToBankAccountId)
	transferCommand := NewTransferCommand(from, to, i.Amount)
	transferCommand.Execute()
	mt.BankAccountRepo.Update(*from)
	mt.BankAccountRepo.Update(*to)
}
