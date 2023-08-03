package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Product{})

	for i := 0; i < 100; i++ {
		hash := md5.Sum([]byte(strconv.Itoa(i)))

		db.Create(&Product{
			Code:  hex.EncodeToString(hash[:2]),
			Price: uint(i * 1000000),
		})
	}

	var r []int
	for i := 0; i < 100; i++ {
		r = append(r, i)
	}
	db.Unscoped().Delete(&Product{}, r)
}
