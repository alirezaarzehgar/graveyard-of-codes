package main

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type X2 struct {
	X int
}

func (x *X2) Scan(v any) (err error) {
	x.X = int(v.(int64))
	return
}

func (x X2) GormDataType() string {
	return "integer"
}

func (x X2) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{SQL: fmt.Sprint(x.X * 2)}
}

type U5 struct {
	X X2
}

func main() {
	db, err := gorm.Open(sqlite.Open("l5.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&U5{}) {
		db.Migrator().DropTable(&U5{})
	}
	db.AutoMigrate(&U5{})

	err = db.Create(&U5{X: X2{10}}).Error
	if err != nil {
		log.Fatal(err)
	}

	var u U5
	db.First(&u)
	fmt.Printf("%#v\n", u)
}
