package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("Não deve criar uma conta com nome inválido", func(t *testing.T) {
		_, err := NewAccount("John", "john.doe@mail.com", "11111111111", "123456", "sha1")
		assert.EqualError(t, err, "invalid name")
	})

	t.Run("Não deve criar uma conta com e-mail inválido", func(t *testing.T) {
		_, err := NewAccount("John Doe", "john.doe@mail", "11111111111", "123456", "sha1")
		assert.EqualError(t, err, "invalid email")
	})

	t.Run("Não deve criar uma conta com documento inválido", func(t *testing.T) {
		_, err := NewAccount("John Doe", "john.doe@mail.com", "1111111", "123456", "sha1")
		assert.EqualError(t, err, "invalid document")
	})

	t.Run("Deve criar senha valida em plainText", func(t *testing.T) {
		account, _ := NewAccount("John Doe", "john.doe@mail.com", "11111111111", "123456", "plaintext")
		password := account.Password.PasswordMatches("123456")
		assert.Equal(t, true, password)
	})

	t.Run("Deve criar senha valida em SHA1", func(t *testing.T) {
		account, _ := NewAccount("John Doe", "john.doe@mail.com", "11111111111", "123456", "sha1")
		password := account.Password.PasswordMatches("123456")
		assert.Equal(t, true, password)
	})
}

func TestAccountDriver(t *testing.T) {
	t.Run("Deve criar conta de motorista", func(t *testing.T) {
		account := &Driver{
			Account: Account{
				Name:     "John Doe",
				Email:    "john.doe@mail.com",
				Document: "11111111111",
				Password: &PasswordSHA1{
					Value: "123456",
				},
			},
			CarPlate: "AAA9999",
		}

		assert.Equal(t, "John Doe", account.Account.Name)
		assert.Equal(t, "john.doe@mail.com", account.Account.Email)
	})

	t.Run("Não deve criar uma conta com a placa do carro inválida", func(t *testing.T) {
		account, _ := NewAccount("John Doe", "john.doe@mail.com", "11111111111", "123456", "sha1")
		_, err := NewDriver(*account, "AAA999")
		assert.EqualError(t, err, "invalid carplate")
	})
}

func TestAccountPassenger(t *testing.T) {
	t.Run("Deve criar uma nova conta para passeiro", func(t *testing.T) {
		account := &Passenger{
			Account: Account{
				Name:     "John Doe",
				Email:    "john.doe@mail.com",
				Document: "11111111111",
				Password: &PasswordSHA1{
					Value: "123456",
				},
			},
			CardHolder: "JOHN DOE",
			CardNumber: "1111 1111 1111 1111",
			ExpDate:    "08/28",
			Cvv:        "123",
		}

		assert.Equal(t, "John Doe", account.Account.Name)
		assert.Equal(t, "john.doe@mail.com", account.Account.Email)
	})

	t.Run("Não deve criar conta de passageiro com CVV inválido", func(t *testing.T) {
		account, _ := NewAccount("John Doe", "john.doe@mail.com", "11111111111", "123456", "sha1")
		_, err := NewPassenger(*account, "JOHN DOE", "1111 1111 1111 1111", "08/28", "12")
		assert.EqualError(t, err, "invalid cvv")
	})
}
