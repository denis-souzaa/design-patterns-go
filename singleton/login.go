package singleton

type Login struct {
	UserRepo UserRepository
}

type InputLogin struct {
	Email    string
	Password string
}

type Output struct {
	Success bool
}

func NewLogin() *Login {
	return &Login{
		UserRepo: Instance(),
	}
}

func (s *Login) Execute(i InputLogin) Output {
	success := false
	user := s.UserRepo.GetByEmail(i.Email)
	if user != nil && user.PasswordMatches(i.Password) {
		success = true
	}

	return Output{Success: success}
}
