package bridge

import (
	"crypto/sha1"
	"errors"
)

type Password interface {
	PasswordMatches(password string) bool
}

type PasswordPlainText struct {
	Value string
}

func NewPasswordPlainText(password string) *PasswordPlainText {
	return &PasswordPlainText{Value: password}
}

func (pp *PasswordPlainText) PasswordMatches(password string) bool {
	return pp.Value == password
}

type PasswordSHA1 struct {
	Value string
}

func NewPasswordSHA1(password string) *PasswordSHA1 {
	value := sha1Hex(password)
	return &PasswordSHA1{Value: value}
}

func (ps *PasswordSHA1) PasswordMatches(password string) bool {
	return ps.Value == sha1Hex(password)
}

func sha1Hex(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	return string(h.Sum(nil))
}

type PasswordFactory struct {
}

func (pf *PasswordFactory) Create(password, passwordType string) (Password, error) {
	if passwordType == "plaintext" {
		return NewPasswordPlainText(password), nil
	}
	if passwordType == "sha1" {
		return NewPasswordSHA1(password), nil
	}
	return nil, errors.New("invalid type passsword")
}
