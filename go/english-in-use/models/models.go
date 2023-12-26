package models

import (
	"fmt"

	"github.com/docker/docker/pkg/homedir"
)

type Word struct {
	ID     uint   `gorm:"primaryKey"`
	Value  string `gorm:"not null; unique"`
	Source string `gorm:"not null"`
}

func GetDbPath() string {
	home, _ := homedir.GetDataHome()
	path := fmt.Sprintf("%s/english-in-use.sqlite", home)
	fmt.Println(path)
	return path
}
