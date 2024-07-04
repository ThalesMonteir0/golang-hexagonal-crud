package controller

import (
	"github.com/ThalesMonteir0/golang-hexagonal-crud/adapter/input/model/request"
	"github.com/ThalesMonteir0/golang-hexagonal-crud/application/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (u *userController) CreateUser(c *gin.Context) {
	var user request.UserRequest

	if err := c.ShouldBindJSON(&user); err != nil {

	}

	userDomain := domain.UserDomain{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Active:   true,
	}

	if err := u.service.CreateUser(userDomain); err != nil {

	}

	c.JSON(http.StatusOK, "ok")
	return
}
