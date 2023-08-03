package main

import (
	"fmt"
	"log"
	"math/rand"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Job struct {
	Title      string
	Experience uint `gorm:"not null; check:job_experience >= 0"`
}

type U2 struct {
	gorm.Model
	Name string `gorm:"not null"`
	Age  uint   `gorm:"check:age > 0;not null"`
	Job  Job    `gorm:"embedded;embeddedPrefix:job_"`
}

func showData(msg string, users []U2) {
	fmt.Println(msg, ":")
	for _, u := range users {
		fmt.Println(u.ID, u.Name, u.Age, u.Job)
	}
	fmt.Println()
}

type gr struct {
	I uint
	X string
}

func main() {
	jobs := []Job{
		{Title: "Back-end Developer", Experience: 25},
		{Title: "Front-end Developer", Experience: 3},
		{Title: "DevOps Engineer", Experience: 0},
		{Title: "SRE", Experience: 15},
		{Title: "Android Developer", Experience: 13},
	}

	names := []string{
		"Ali", "Reza", "Sadegh", "Mohammad",
		"Hamid", "Amir", "Sadra", "Javad",
		"Saleh", "Pooya", "Morteza", "Mostafa", "Milad"}

	db, err := gorm.Open(sqlite.Open("l2.db"))
	if err != nil {
		log.Fatal(err)
	}

	if db.Migrator().HasTable(&U2{}) {
		db.Migrator().DropTable(&U2{})
	}
	db.AutoMigrate(&U2{})

	for _, name := range names {
		for _, job := range jobs {
			err = db.Create(&U2{
				Name: name,
				Age:  uint(rand.Intn(60) + 10),
				Job:  job,
			}).Error
			if err != nil {
				log.Println(err)
			}
		}
	}

	var users []U2

	db.Where("UPPER(name) <> ? ", "ALI").Find(&users)
	showData("Get all mached records", users)

	db.Where("UPPER(name) IN ?", []string{"ALI", "REZA"}).Find(&users)
	showData("Get ali and reza", users)

	// db.Where(&U2{Job: jobs[3]}, &U2{Name: "Ali"}).Find(&users)
	db.Where(&U2{Name: "Ali", Job: jobs[3]}).Find(&users)
	showData("Get Ali SRE", users)

	db.Where([]int64{3, 4, 5}).Find(&users)
	showData("Get 3, 4, 5", users)

	db.Where(map[string]any{"job_experience": 0}).Find(&users)
	showData("Get zero experiences", users)

	s := U2{
		Name: "Mohammad",
		Job:  jobs[0],
	}
	db.Where(&s, "job_experience").Find(&users)
	showData("Get all n experienced", users)

	db.Find(&users, "name <> ? and age > ?", "Ali", 65)
	showData("Get Ali with 65<=age", users)

	db.Not("name <> ?", "Ali").Find(&users)
	showData("Get alis", users)

	db.Not("name <> ?", "Mohmmad").Where("age < 20").Or("age > 30").Where("age < 40").Find(&users)
	showData("or condition", users)

	db.Where("name = ?", "Ali").Order("age desc, name").Find(&users)
	showData("order by age", users)

	db.Limit(3).Find(&users)
	showData("limit 3", users)

	db.Limit(3).Find(&users).Limit(-1).Offset(5).Limit(3).Find(&users)
	showData("limit -1", users)

	var r []gr
	db.Model(&U2{}).Select("COUNT(id) as I, name as X").Group("name").Find(&r)
	fmt.Println("group by job_title", r)

	db.Model(&U2{}).Select("COUNT(id) as I, job_title as X").Group("job_title").Having("age > ?", 20).Find(&r)
	fmt.Println("group by job_title having age > n", r)

	db.Distinct("name").Order("age").Find(&users)
	showData("distinct names", users)

	db.Raw("SELECT sqlite_version() as X").Scan(&r)
	fmt.Println(r)
}
