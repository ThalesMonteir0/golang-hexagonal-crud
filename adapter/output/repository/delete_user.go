package repository

import "github.com/ThalesMonteir0/golang-hexagonal-crud/application/domain"

func (ur *userRepository) DeleteUser(userDomain domain.UserDomain) error {
	return nil
}
