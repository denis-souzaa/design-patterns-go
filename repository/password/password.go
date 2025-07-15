package password

import "errors"

type Password struct {
	value string
}

func New(password string) (*Password, error) {
	if len(password) < 8 {
		return nil, errors.New("minimum lenght 8 caracteres")
	}
	return &Password{value: password}, nil
}

func (p *Password) Value() string {
	return p.value
}
