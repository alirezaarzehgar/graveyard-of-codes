package model

type User struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Posts    []Post
}
