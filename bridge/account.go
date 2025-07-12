package bridge

import (
	"errors"
	"regexp"
)

type Account struct {
	Name     string
	Email    string
	Document string
	Password Password
}

func NewAccount(name, email, document, password, passwordType string) (*Account, error) {
	validName, err := regexp.MatchString(".+ .+", name)
	if err != nil {
		panic(err)
	}

	if !validName {
		return nil, errors.New("invalid name")
	}
	var emailRE = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	validEmail := emailRE.MatchString(email)

	if !validEmail {
		return nil, errors.New("invalid email")
	}

	if len(document) != 11 {
		return nil, errors.New("invalid document")
	}

	passFactory := &PasswordFactory{}
	pass, err := passFactory.Create(password, passwordType)
	if err != nil {
		panic(err)
	}
	return &Account{Name: name, Email: email, Document: document, Password: pass}, nil
}
