package model

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"unique; not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	Price       uint   `gorm:"not null" json:"price"`
	Rate        uint   `json:"rate"`
	OwnerID     uint
	Owner       User
}

func CreateProduct(p *Product) error {
	err := DB.Create(p).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return echo.ErrConflict
	}
	if err != nil {
		return err
	}
	return nil
}

func GetProductByID(id uint) (*Product, error) {
	var p Product
	var mysqlErr *mysql.MySQLError
	err := DB.Preload("Owner").First(&p, "id = ?", id).Error
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return nil, echo.ErrConflict
	}
	if err != nil {
		return nil, err
	}
	p.Owner.Password = ""
	return &p, nil
}

func EditProductByID(id uint, p *Product) error {
	p.ID = id
	err := DB.Model(p).Select("name", "description", "price").Updates(p).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteProductByID(id uint) error {
	if err := DB.Unscoped().Delete(&Product{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
