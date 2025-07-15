package email

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func New(email string) (*Email, error) {
	if valid, err := regexp.MatchString(".+@.+.com", email); err != nil {
		return nil, err
	} else if !valid {
		return nil, errors.New("invalid email")
	}

	return &Email{value: email}, nil
}

func (e *Email) Value() string {
	return e.value
}
