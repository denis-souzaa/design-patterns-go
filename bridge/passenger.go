package bridge

import "errors"

type Passenger struct {
	Account    Account
	CardHolder string
	CardNumber string
	ExpDate    string
	Cvv        string
}

func NewPassenger(account Account, cardHolder, cardNumber, expDate, cvv string) (*Passenger, error) {
	if len(cvv) != 3 {
		return nil, errors.New("invalid cvv")
	}

	return &Passenger{Account: account, CardHolder: cardHolder, CardNumber: cardNumber, ExpDate: expDate, Cvv: cvv}, nil
}
