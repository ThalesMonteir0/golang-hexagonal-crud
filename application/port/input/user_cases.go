package input

import "github.com/ThalesMonteir0/golang-hexagonal-crud/application/domain"

type UserServiceInterface interface {
	FindUser(userDomain domain.UserDomain) (*domain.UserDomain, error)
	FindUserByEmail(string) (*domain.UserDomain, error)
	DeleteUser(userDomain domain.UserDomain) error
	CreateUser(userDomain domain.UserDomain) error
}
