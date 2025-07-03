package main

type LoanFactory interface {
	CreateLoan(amount, income float64, installments int) Loan
	CreateInstallmentCalculator() InstallmentCalculator
}

type MortgageFactory struct {
}

func NewMortgageFactory() *MortgageFactory {
	return &MortgageFactory{}
}

func (m *MortgageFactory) CreateLoan(amount, income float64, installments int) Loan {
	ml, _ := NewMortgageLoan(amount, income, installments)
	return ml
}

func (m *MortgageFactory) CreateInstallmentCalculator() InstallmentCalculator {
	return NewSACInstallmentCalculator()
}

type CarFactory struct {
}

func (c *CarFactory) CreateLoan(amount, income float64, installments int) *CarLoan {
	cl, err := NewCarLoan(amount, income, installments)
	if err != nil {
		return nil
	}
	return cl
}

func (c *CarFactory) CreateInstallmentCalculator() InstallmentCalculator {
	return NewPriceInstallmentCalculator()
}
