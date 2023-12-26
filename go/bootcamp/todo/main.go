package main

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	ID       uint
	Username string
	Password string
}

var scanner = bufio.NewScanner(os.Stdin)
var userStorage []User

func main() {
	for {
		fmt.Printf("Enter command: ")
		scanner.Scan()
		runCommand(scanner.Text())
	}
}

func runCommand(command string) {
	switch command {
	case "register":
		registerUser()
	case "login":
		loginUser()
	case "create-category":
		createCategory()
	case "create-task":
		createTask()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("Invalid command:", command)
	}
}

func registerUser() {
	id := uint(len(userStorage) + 1)

	fmt.Printf("Enter username: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	userStorage = append(userStorage, User{id, name, password})
}

func loginUser() {
	fmt.Printf("Enter username: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Printf("Enter password: ")
	scanner.Scan()
	password := scanner.Text()

	for _, user := range userStorage {
		if user.Username == name {
			if user.Password == password {
				fmt.Println("Logged in!")
				return
			}
			break
		}
	} 
	fmt.Println("Invalid password!")
}

func createCategory() {
	fmt.Printf("Enter category title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Printf("Enter category color: ")
	scanner.Scan()
	color := scanner.Text()

	fmt.Print("category: ", title, color)
}

func createTask() {
	fmt.Printf("Enter task title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Printf("Enter task category: ")
	scanner.Scan()
	category := scanner.Text()

	fmt.Printf("Enter task due date: ")
	scanner.Scan()
	dueDate := scanner.Text()

	fmt.Println("task: ", title, category, dueDate)
}
