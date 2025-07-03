package main

type RepositoryFactory interface {
	CreateLoanRepository() LoanRepository
	CreateInstallmentRepository() InstallmentRepository
}

type RepositoryMemoryFactory struct {
}

func (rm *RepositoryMemoryFactory) CreateLoanRepository() LoanRepository {
	l := LoanRepositoryMemory{}
	return l.GetInstance()
}

func (rm *RepositoryMemoryFactory) CreateInstallmentRepository() InstallmentRepository {
	i := InstallmentReposityMemory{}
	return i.GetInstance()
}

func NewRepositoryMemoryFactory() *RepositoryMemoryFactory {
	return &RepositoryMemoryFactory{}
}
