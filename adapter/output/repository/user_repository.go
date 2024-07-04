package repository

import (
	"github.com/ThalesMonteir0/golang-hexagonal-crud/application/port/output"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) output.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}
