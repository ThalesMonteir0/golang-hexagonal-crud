package controller

import (
	"github.com/ThalesMonteir0/golang-hexagonal-crud/application/port/input"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service input.UserServiceInterface
}

type UserControllerInterface interface {
	GetUser(*gin.Context)
	CreateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

func NewUserController(service input.UserServiceInterface) UserControllerInterface {
	return &userController{
		service: service,
	}
}
