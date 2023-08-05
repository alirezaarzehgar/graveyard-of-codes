package model

import "gorm.io/gorm"

type Joke struct {
	gorm.Model
	Content  string
	Author   string
	Rate     uint
	Comments []Comment
}
