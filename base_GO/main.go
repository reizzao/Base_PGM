package main

import (
	"githubendereco/usecases/CreateUser"
	"githubendereco/usecases/ListUsers"
)

func main() {
	// fmt.Println("Hello Main")
	CreateUser.CreateUserTester()
	ListUsers.ListUserTester()
}
