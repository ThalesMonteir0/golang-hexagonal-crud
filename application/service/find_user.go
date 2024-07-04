package service

import "github.com/ThalesMonteir0/golang-hexagonal-crud/application/domain"

func (us *userService) FindUser(userDomain domain.UserDomain) (*domain.UserDomain, error) {
	return nil, nil
}

func (us *userService) FindUserByEmail(email string) (*domain.UserDomain, error) {
	return nil, nil
}
