package main

import "sync"

var onceI sync.Once
var singleInstanceI *InstallmentReposityMemory

type InstallmentRepository interface {
	Save(installment Installment)
	ListByLoanId(loanId string) []Installment
}

type InstallmentReposityMemory struct {
	Installments []Installment
}

func (i *InstallmentReposityMemory) GetInstance() *InstallmentReposityMemory {
	if singleInstanceI == nil {
		onceI.Do(func() {
			singleInstanceI = &InstallmentReposityMemory{}
		})
	}
	return singleInstanceI
}

func NewInstallmentRepositoryInMemory() *InstallmentReposityMemory {
	return &InstallmentReposityMemory{}
}

func (m *InstallmentReposityMemory) Save(i Installment) {
	m.Installments = append(m.Installments, i)
}

func (i *InstallmentReposityMemory) ListByLoanId(loanId string) []Installment {
	installments := []Installment{}
	for _, i := range i.Installments {
		if i.LoanId == loanId {
			installments = append(installments, i)
		}
	}

	return installments
}
