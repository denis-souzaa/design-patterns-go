package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustCreateNewUser(t *testing.T) {
	u, err := New("John Doe", "john.doe@mail.com", "abc123456")
	assert.Equal(t, nil, err)
	assert.Equal(t, "John Doe", u.Name())
	assert.Equal(t, "john.doe@mail.com", u.Email())
	assert.Equal(t, "abc123456", u.Password())
	assert.Equal(t, "active", u.Status())
}

func TestMustChangePasswordUser(t *testing.T) {
	u, err := New("John Doe", "john.doe@mail.com", "abc123456")
	assert.Equal(t, nil, err)
	u.UpdatePassword("asf45679")
	assert.Equal(t, "asf45679", u.Password())
}

func TestNotMustChangePasswordUserWhenLessEightCaracteres(t *testing.T) {
	u, _ := New("John Doe", "john.doe@mail.com", "abc123456")
	err := u.UpdatePassword("asf")
	assert.EqualError(t, err, "minimum lenght 8 caracteres")
}

func TestNotMustChangeEmailUserWhenInvalid(t *testing.T) {
	u, _ := New("John Doe", "john.doe@mail.com", "abc123456")
	err := u.UpdateEmail("john.doe")
	assert.EqualError(t, err, "invalid email")
}

func TestMustBlockUser(t *testing.T) {
	u, err := New("John Doe", "john.doe@mail.com", "abc123456")
	assert.Equal(t, nil, err)
	err = u.Block()
	assert.Equal(t, nil, err)
	assert.Equal(t, string(Blocked), u.Status())
}
