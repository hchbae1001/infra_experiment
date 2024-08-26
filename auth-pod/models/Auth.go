package models

type Auth struct {
	AuthID       int    `json:"auth_id"`
	UserID       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
}
