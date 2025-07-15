package user

import (
	"denis-souzaa/design-patterns-go/config"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMustSaveNewUser(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src).Intn(1000)
	db, err := config.New()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	userRepo := NewRepositoryDatabase(db)
	email := fmt.Sprintf("john.doe%d@mail.com", r)
	user, err := New("John Doe", email, "abc123456")
	if err != nil {
		log.Fatal(err)
	}
	userRepo.Save(*user)
	savedUser, err := userRepo.ByEmail(user.Email())
	assert.Nil(t, nil, err)
	assert.Equal(t, "John Doe", savedUser.Name())
	assert.Equal(t, email, savedUser.Email())
	assert.Equal(t, "abc123456", savedUser.Password())
	assert.Equal(t, string(Active), savedUser.Status())
	userRepo.Delete(email)
}

func TestMustUpdateUser(t *testing.T) {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src).Intn(1000)
	db, err := config.New()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	userRepo := NewRepositoryDatabase(db)
	email := fmt.Sprintf("john.doe%d@mail.com", r)
	user, err := New("John Doe", email, "abc123456")
	if err != nil {
		log.Fatal(err)
	}
	userRepo.Save(*user)
	savedUser, err := userRepo.ByEmail(user.Email())
	assert.Nil(t, nil, err)
	savedUser.UpdatePassword("asd45678")
	userRepo.Update(*savedUser)
	updatedUser, err := userRepo.ByEmail(user.Email())
	assert.Nil(t, nil, err)
	assert.Equal(t, "asd45678", updatedUser.Password())
	userRepo.Delete(email)
}

func TestMustListUsers(t *testing.T) {
	db, err := config.New()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	userRepo := NewRepositoryDatabase(db)
	user1, err := New("John Doe", "john.doe1@mail.com", "abc123456")
	user2, _ := New("John Doe", "john.doe2@mail.com", "abc123456")
	user3, _ := New("John Doe", "john.doe3@mail.com", "abc123456")
	if err != nil {
		log.Fatal(err)
	}
	userRepo.Save(*user1)
	userRepo.Save(*user2)
	userRepo.Save(*user3)
	users, err := userRepo.List()
	assert.Nil(t, nil, err)
	assert.Len(t, users, 3)
	userRepo.Delete("john.doe1@mail.com")
	userRepo.Delete("john.doe2@mail.com")
	userRepo.Delete("john.doe3@mail.com")
}
