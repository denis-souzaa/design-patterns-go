package command

type TypeTransaction string

const (
	Debit  TypeTransaction = "debit"
	Credit TypeTransaction = "credit"
)

type Transaction struct {
	Type   TypeTransaction
	Amount float64
}
