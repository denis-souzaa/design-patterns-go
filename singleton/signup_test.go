package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSigup(t *testing.T) {
	signup := NewSignup()
	login := NewLogin()
	inputSignup := InputSignup{Name: "John Doe", Email: "john.doe@mail.com", Password: "123456"}
	signup.Execute(inputSignup)
	inputLogin := InputLogin{Email: "john.doe@mail.com", Password: "123456"}
	outputLogin := login.Execute(inputLogin)
	assert.Equal(t, true, outputLogin.Success)
}
