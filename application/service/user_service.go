package service

import (
	"github.com/ThalesMonteir0/golang-hexagonal-crud/application/port/input"
	"github.com/ThalesMonteir0/golang-hexagonal-crud/application/port/output"
)

type userService struct {
	repository output.UserRepositoryInterface
}

func NewUserService(repository output.UserRepositoryInterface) input.UserServiceInterface {
	return &userService{
		repository: repository,
	}
}
