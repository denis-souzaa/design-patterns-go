package singleton

import (
	"github.com/google/uuid"
)

type User struct {
	UserId   string
	Name     string
	Email    string
	Password string
}

func NewUser(userId, name, email, password string) *User {
	return &User{UserId: userId, Name: name, Email: email, Password: password}
}

func (u *User) Create(name, email, password string) *User {
	userId := uuid.New().String()
	return &User{UserId: userId, Name: name, Email: email, Password: password}
}

func (u *User) PasswordMatches(p string) bool {
	return u.Password == p
}
