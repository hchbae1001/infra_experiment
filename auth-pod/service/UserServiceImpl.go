package service

import (
	"auth-pod/models"
	"auth-pod/repository"
	"auth-pod/util"
	"errors"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(user *models.User) (int64, error) {

	existingUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return 0, err
	}
	if existingUser != nil {
		return 0, errors.New("email already in use")
	}

	encryptedPassword, err := util.Encrypt(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = encryptedPassword

	userID, err := s.userRepo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *userService) GetUserByEmail(email string) (*models.User, error) {
	return s.userRepo.GetUserByEmail(email)
}

func (s *userService) GetUserById(userId int) (*models.User, error) {
	return s.userRepo.GetUserById(userId)
}
