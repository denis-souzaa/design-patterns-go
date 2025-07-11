package singleton

import "sync"

var once = &sync.Once{}
var instance *UserRepositoryMemory

type UserRepository interface {
	Save(u User)
	GetByEmail(email string) *User
}

type UserRepositoryMemory struct {
	users []User
}

func Instance() *UserRepositoryMemory {
	once.Do(func() {
		instance = &UserRepositoryMemory{}
	})

	return instance
}

func (um *UserRepositoryMemory) Save(u User) {
	um.users = append(um.users, u)
}

func (um *UserRepositoryMemory) GetByEmail(email string) *User {
	for _, u := range um.users {
		if u.Email == email {
			return &u
		}
	}
	return nil
}
