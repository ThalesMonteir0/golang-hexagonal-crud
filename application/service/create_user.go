package service

import (
	"errors"
	"github.com/ThalesMonteir0/golang-hexagonal-crud/application/domain"
)

func (us *userService) CreateUser(userDomain domain.UserDomain) error {
	if user, _ := us.FindUserByEmail(userDomain.Email); user != nil {
		return errors.New("e-mail already registered")
	}

	if err := userDomain.EncryptPassword(); err != nil {
		return errors.New("unable encrypted password")
	}

	if err := us.repository.CreateUser(userDomain); err != nil {
		return errors.New("unable created user")
	}

	return nil
}
