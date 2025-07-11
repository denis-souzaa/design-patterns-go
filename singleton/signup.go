package singleton

import "fmt"

type Signup struct {
	UserRepo UserRepository
}

type InputSignup struct {
	Name     string
	Email    string
	Password string
}

func NewSignup() *Signup {
	return &Signup{
		UserRepo: Instance(),
	}
}

func (s *Signup) Execute(i InputSignup) {
	user := &User{}
	user = user.Create(i.Name, i.Email, i.Password)
	fmt.Println(user)
	s.UserRepo.Save(*user)
}
