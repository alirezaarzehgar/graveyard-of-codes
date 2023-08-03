package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreditCard struct {
	gorm.Model
	Number string
	U6ID   uint
}

type U6 struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}

func main() {
	db, err := gorm.Open(sqlite.Open("l6.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&CreditCard{}) {
		db.Migrator().DropTable(&U6{}, &CreditCard{})
	}
	db.AutoMigrate(&CreditCard{}, &U6{})

	u := U6{
		Name:       "Alireza",
		CreditCard: CreditCard{Number: "AASDASD123ASZXX"},
	}

	db.Create(&u)
	u.ID++
	db.Omit("CreditCard").Create(&u)
	u.ID++
	db.Select("Name").Create(&u)
	u.ID++
	db.Omit(clause.Associations).Create(&u)
}
