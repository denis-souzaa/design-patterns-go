package main

import (
	"errors"
	"sync"
)

var once sync.Once
var singleInstance *LoanRepositoryMemory

type LoanRepository interface {
	Save(loan Loan)
	GetById(loanId string) (Loan, error)
}

type LoanRepositoryMemory struct {
	Loan []Loan
}

func (m *LoanRepositoryMemory) GetInstance() *LoanRepositoryMemory {
	if singleInstance == nil {
		once.Do(
			func() {
				singleInstance = &LoanRepositoryMemory{}
			})
	}
	return singleInstance
}

func (m *LoanRepositoryMemory) Save(loan Loan) {
	m.Loan = append(m.Loan, loan)
}

func (m *LoanRepositoryMemory) GetById(loanId string) (Loan, error) {
	for _, l := range m.Loan {
		if l.GetLoanId() == loanId {
			return l, nil
		}
	}

	return nil, errors.New("Loan not found")
}
