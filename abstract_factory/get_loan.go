package main

type GetLoan struct {
	LoanRepo        LoanRepository
	InstallmentRepo InstallmentRepository
}

func NewGetLoan(rf RepositoryFactory) *GetLoan {
	return &GetLoan{LoanRepo: rf.CreateLoanRepository(), InstallmentRepo: rf.CreateInstallmentRepository()}
}

type InputLoan struct {
	LoanId string
}

type OutputLoan struct {
	Amount       float64
	Income       float64
	Installments []InstallmentOutput
}

type InstallmentOutput struct {
	Number       int
	Amount       float64
	Amortization float64
	Interest     float64
	Balance      float64
}

func (l *GetLoan) Execute(i InputLoan) (OutputLoan, error) {
	loan, err := l.LoanRepo.GetById(i.LoanId)
	if err != nil {
		return OutputLoan{}, err
	}
	installments := l.InstallmentRepo.ListByLoanId(i.LoanId)
	output := OutputLoan{Amount: loan.GetAmount(), Income: loan.GetIncome(), Installments: []InstallmentOutput{}}
	for _, i := range installments {
		iout := InstallmentOutput{
			Number:       i.Number,
			Amount:       i.Amount,
			Amortization: i.Amortization,
			Interest:     i.Interest,
			Balance:      i.Balance}
		output.Installments = append(output.Installments, iout)
	}

	return output, nil
}
