package model

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("MYSQL_DATABASE"))
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	sqlDB, _ := DB.DB()
	if err = sqlDB.Ping(); err != nil {
		return err
	}

	if err := DB.AutoMigrate(&User{}, &Product{}); err != nil {
		return err
	}

	return nil
}
