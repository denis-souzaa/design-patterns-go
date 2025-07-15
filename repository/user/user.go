package user

import (
	"denis-souzaa/design-patterns-go/repository/email"
	"denis-souzaa/design-patterns-go/repository/password"
	"errors"
	"log"
)

type Status string

const (
	Active  Status = "active"
	Blocked Status = "blocked"
)

type User struct {
	name     string
	email    email.Email
	password password.Password
	status   Status
}

func New(name, mail, pass string) (*User, error) {
	password, err := password.New(pass)
	if err != nil {
		return nil, err
	}
	email, err := email.New(mail)
	if err != nil {
		return nil, err
	}

	return &User{name: name, email: *email, password: *password, status: Active}, nil
}

func new(name, mail, pass, status string) *User {
	password, err := password.New(pass)
	if err != nil {
		log.Fatal(err)
	}
	email, err := email.New(mail)
	if err != nil {
		log.Fatal(err)
	}

	return &User{name: name, email: *email, password: *password, status: Status(status)}
}

func (u *User) UpdatePassword(pass string) error {
	password, err := password.New(pass)
	if err != nil {
		return err
	}
	u.password = *password
	return nil
}

func (u *User) UpdateEmail(mail string) error {
	email, err := email.New(mail)
	if err != nil {
		return err
	}
	u.email = *email
	return nil
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email.Value()
}

func (u *User) Password() string {
	return u.password.Value()
}

func (u *User) Status() string {
	return string(u.status)
}

func (u *User) Block() error {
	if u.Status() == string(Blocked) {
		return errors.New("user already blocked")
	}
	u.status = Blocked
	return nil
}
