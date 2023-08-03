package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	FirstName string `gorm:"not null"`
	LasstName string `gorm:"not null"`
	JobCode   int
	Job       Job
}

type Job struct {
	Code      int    `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	CompanyID int
	Company   Company
}

type Company struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`
	Phone   string `gorm:"not null"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&Company{}, &Job{}, &User{})
	}
	db.AutoMigrate(&Job{}, &Company{}, &User{})

	db.Create(&User{
		FirstName: "Alireza",
		LasstName: "Arzehar",
		Job: Job{
			Title: "Software Developer",
			Company: Company{
				Name:    "Zharfpouyan Toos",
				Address: "Mashhad, Emamat 3",
				Phone:   "09155093114",
			},
		},
	})

	var u User
	db.Preload("Job.Company").Find(&u)
	fmt.Printf("%#v\n", u)

	err = db.Delete(&u).Error
	if err != nil {
		log.Fatal(err)
	}

	u = User{}
	db.Preload("Job.Company").Find(&u)
	fmt.Printf("%#v\n", u)

}
