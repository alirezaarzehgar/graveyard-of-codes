package model

type Post struct {
	ID      int    `gorm:"primaryKey"`
	Title   string `gorm:"not null;unique"`
	Content string `gorm:"not null"`
	UserID  int
	User    User
}
