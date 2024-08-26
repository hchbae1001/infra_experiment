package service

import (
	"auth-pod/models"
)

type UserService interface {
	CreateUser(user *models.User) (int64, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(userId int) (*models.User, error)
}
