package main

type ApplyForLoan struct {
	LoanRepo        LoanRepository
	InstallmentRepo InstallmentRepository
	LoanFactory     LoanFactory
}

type Input struct {
	Amount      float64
	Income      float64
	Installment int
}

type Output struct {
	LoanId string
}

func NewApplyForLoan(rf RepositoryFactory, lf LoanFactory) ApplyForLoan {
	return ApplyForLoan{
		LoanRepo:        rf.CreateLoanRepository(),
		InstallmentRepo: rf.CreateInstallmentRepository(),
		LoanFactory:     lf,
	}
}

func (afl *ApplyForLoan) Execute(i Input) (Output, error) {

	loan := afl.LoanFactory.CreateLoan(i.Amount, i.Income, i.Installment)
	installmentCalculator := afl.LoanFactory.CreateInstallmentCalculator()
	installments := installmentCalculator.calculate(loan)
	afl.LoanRepo.Save(loan)
	for _, i := range installments {
		afl.InstallmentRepo.Save(i)
	}
	return Output{LoanId: loan.GetLoanId()}, nil
}
