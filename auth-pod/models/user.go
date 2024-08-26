package models

type User struct {
	UserID   int    `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
