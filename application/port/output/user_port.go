package output

import "github.com/ThalesMonteir0/golang-hexagonal-crud/application/domain"

type UserRepositoryInterface interface {
	FindUser(userDomain domain.UserDomain) (*domain.UserDomain, error)
	CreateUser(userDomain domain.UserDomain) error
	DeleteUser(userDomain domain.UserDomain) error
}
