package bridge

import (
	"errors"
	"regexp"
)

type Driver struct {
	Account  Account
	CarPlate string
}

func NewDriver(account Account, carPlate string) (*Driver, error) {
	validCarPlate, err := regexp.MatchString("[A-Z]{3}[0-9]{}", carPlate)
	if err != nil {
		panic(err)
	}
	if !validCarPlate {
		return nil, errors.New("invalid carplate")
	}

	return &Driver{Account: account, CarPlate: carPlate}, nil
}
