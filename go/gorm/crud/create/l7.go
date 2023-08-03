package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Counter struct {
	Word  string `gorm:"primaryKey"`
	Count uint   `gorm:"default:1"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("l7.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&Counter{}) {
		db.Migrator().DropTable(&Counter{})
	}
	db.AutoMigrate(&Counter{})

	keys := []string{"a", "b", "c", "d"}

	rnd := rand.New(rand.NewSource(time.Now().UnixMilli()))
	for _, key := range keys {
		r := rnd.Intn(10) + 3
		for i := 0; i < r; i++ {
			db.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "word"}},
				DoUpdates: clause.Assignments(map[string]any{"count": gorm.Expr("count+1")}),
			}).Create(&Counter{Word: key})
		}
	}

	var cs []Counter
	db.Find(&cs)
	for _, v := range cs {
		fmt.Println(v)
	}
}
