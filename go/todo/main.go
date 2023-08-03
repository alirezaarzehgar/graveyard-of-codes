package main

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/pterm/pterm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const MAX_HEIGHT = 20

var db *gorm.DB

type User struct {
	ID       uint
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Loggedin bool   `gorm:"not null;default:false"`
	Tasks    []Task
}

type Task struct {
	ID          uint
	Title       string `gorm:"not null"`
	Description string
	UserID      uint
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func currentUser() *User {
	var user User
	err := db.Where("loggedin = ?", true).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func auth() (*User, error) {
	action, _ := pterm.DefaultInteractiveSelect.
		WithFilter(false).
		WithOptions([]string{"Login", "Register"}).
		Show()

	userStr, _ := pterm.DefaultInteractiveTextInput.Show("Enter username")
	passStr, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Enter password")
	passStr = fmt.Sprintf("%x", sha512.Sum512([]byte(passStr)))
	user := &User{Username: userStr, Password: passStr}

	var errMsg error
	switch action {
	case "Login":
		errMsg = db.Where(&user).First(&user).Error
		user.Loggedin = true
		if errMsg == nil {
			db.Save(&user)
		}
	case "Register":
		user.Loggedin = true
		errMsg = db.Create(user).Error
	}

	return user, errMsg
}

func manageTasks(user *User) {
	db.Preload("Tasks").Where(user).First(user)

	var taskTitles []string
	for _, task := range user.Tasks {
		taskTitles = append(taskTitles, task.Title)
	}

	dones, _ := pterm.DefaultInteractiveMultiselect.
		WithOptions(taskTitles).
		Show()

	r := db.Delete(&Task{}, "title in (?)", dones)
	if r.Error != nil {
		pterm.Error.Println(r.Error)
	}
	pterm.Success.Printfln("Done %d tasks", r.RowsAffected)
}

func showTasks(user *User, showTodos bool) {
	db.Preload("Tasks").Where(user).First(user)
	if !showTodos {
		db.Unscoped().Where("deleted_at IS NOT NULL").Find(&user.Tasks)
	}

	var taskTitles []string = []string{"Back"}

	for _, task := range user.Tasks {
		taskTitles = append(taskTitles, task.Title)
	}

	taskTitle := ""
	for {
		taskTitle, _ = pterm.DefaultInteractiveSelect.
			WithOptions(taskTitles).
			WithDefaultOption(taskTitle).
			WithMaxHeight(MAX_HEIGHT).
			Show()

		if taskTitle == "Back" {
			break
		}

		for _, task := range user.Tasks {
			if task.Title == taskTitle {
				pterm.DefaultParagraph.Println(task.Description)
			}
		}
	}
}

func addTask(user *User) {
	taskTitle, _ := pterm.DefaultInteractiveTextInput.Show("Task title")
	taskDescription, _ := pterm.DefaultInteractiveTextInput.
		WithMultiLine().
		Show("Task description")
	_, _ = taskTitle, taskDescription

	if taskTitle == "" {
		pterm.Error.Println("Task title shouldn't be empty")
	}

	err := db.Create(&Task{Title: taskTitle, Description: taskDescription, UserID: user.ID}).Error
	if err != nil {
		pterm.Error.Println(err)
	}
	pterm.Success.Printfln("Task \"%s\" created successfully", taskTitle)
}

func logout(user *User) {
	user.Loggedin = false
	db.Save(&user)
}

func main() {
	var err error
	home, _ := os.UserHomeDir()
	db, err = gorm.Open(sqlite.Open(home+"/.todo.sqlite"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{}, &Task{})

	user := currentUser()
	for user == nil {
		var err error
		user, err = auth()
		if err != nil {
			pterm.Error.Println(err)
			user = nil
		}
	}

	action := ""
	for {
		tasks := []string{"Manage tasks", "Show Todos", "Show Dones", "Add task", "Logout", "Exit"}
		action, _ = pterm.DefaultInteractiveSelect.
			WithFilter(false).
			WithOptions(tasks).
			WithDefaultOption(action).
			WithMaxHeight(len(tasks)).
			Show("Select")

		switch action {
		case "Manage tasks":
			manageTasks(user)
		case "Show Todos":
			showTasks(user, true)
		case "Show Dones":
			showTasks(user, false)
		case "Add task":
			addTask(user)
		case "Logout":
			logout(user)
			os.Exit(0)
		case "Exit":
			os.Exit(0)
		}
	}
}
