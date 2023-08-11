package models

type User struct {
	ID       uint   `gorm:"index"`
	Username string `json:"user"`
	Password string `json:"pass"`
	Email    string `json:"email"`
}
