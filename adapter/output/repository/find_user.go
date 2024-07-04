package repository

import "github.com/ThalesMonteir0/golang-hexagonal-crud/application/domain"

func (ur *userRepository) FindUser(userDomain domain.UserDomain) (*domain.UserDomain, error) {
	return nil, nil
}
