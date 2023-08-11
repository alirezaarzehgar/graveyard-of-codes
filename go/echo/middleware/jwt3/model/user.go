package model

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"user" gorm:"unique; not null"`
	Password string `json:"pass" gorm:"not null"`
}

func hashPassword(pass string) string {
	textPassword := pass
	hashedPassword := sha256.Sum256([]byte(textPassword))
	pass = hex.EncodeToString(hashedPassword[:])
	return pass
}

func RegisterUser(u *User) error {
	u.Password = hashPassword(u.Password)

	err := DB.Create(&u).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return echo.ErrConflict
	}
	if err != nil {
		return err
	}
	return nil
}

func LoginUser(u *User) error {
	u.Password = hashPassword(u.Password)
	r := DB.Where(u).First(u)
	if errors.Is(r.Error, gorm.ErrRecordNotFound) {
		return echo.ErrNotFound
	}
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func DeleteUser(u *User) error {
	return nil
}

func GetUserByID(id uint) (*User, error) {
	return nil, nil
}
