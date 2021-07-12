package domain

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	Base
	Name     string
	Email    string
	Password string
	Token    string
}

func (user *User) Prepare() error {
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error during the password generation: %v", err)
		return err
	}

	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	err = user.validate()

	if err != nil {
		log.Fatalf("Error during the user validation: %v", err)
		return err
	}

	return nil
}

func (user *User) validate() error  {
	return nil
}