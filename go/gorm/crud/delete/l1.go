package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

type Info struct {
	gorm.Model
	Data string
}

func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println(u.Name, "was deleted")
	return
}

func main() {
	db, err := gorm.Open(sqlite.Open("l1.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&User{}) {
		db.Migrator().DropTable(&Info{}, &User{})
	}
	db.AutoMigrate(&User{}, &Info{})

	for i := 0; i < 100; i++ {
		db.Create(&User{Name: fmt.Sprint("U", i+1)})
		db.Create(&Info{Data: fmt.Sprint("I", i+1)})
	}

	var u User
	db.First(&u)
	db.Delete(u)

	db.Where("id < ?", 3).Delete(&User{})

	var us []User
	db.Clauses(clause.Returning{}).Delete(&us, []int{5, 6, 7, 8})
	fmt.Println(us)

	var infs []Info
	// db.Clauses(clause.Returning{}).Where("1 = 1").Delete(&infs)
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).
		Clauses(clause.Returning{}).
		Delete(&infs)
	fmt.Println(infs[:3])

	err = db.Unscoped().Model(&Info{}).Where(1).Update("deleted_at", nil).Error
	if err != nil {
		log.Fatal(err)
	}
	db.Find(&infs)
	fmt.Println(infs)
}
