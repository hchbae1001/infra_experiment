package repository

import (
	"auth-pod/models"
	"database/sql"
	"errors"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(userId int) (*models.User, error)
	CreateUser(user *models.User) (int64, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) GetUserByEmail(email string) (*models.User, error) {
	query := "select * from users where user_id=?"
	row := u.db.QueryRow(query, email)
	var user models.User
	err := row.Scan(&user.UserID, &user.Name, &user.Password, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (u userRepository) GetUserById(userId int) (*models.User, error) {
	query := "select * from users where user_id=?"
	row := u.db.QueryRow(query, userId)
	var user models.User
	err := row.Scan(&user.UserID, &user.Name, &user.Password, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (u userRepository) CreateUser(user *models.User) (int64, error) {

	query := "INSERT INTO users (name, password, email) VALUES (?, ?, ?)"
	result, err := u.db.Exec(query, user.Name, user.Password, user.Email)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, nil
}
