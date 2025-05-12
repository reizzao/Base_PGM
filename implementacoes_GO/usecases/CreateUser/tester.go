package CreateUser

import (
	"fmt"
	"githubendereco/seed"
	"githubendereco/usecases/ListUsers"
)

func CreateUserTester() {
	fmt.Println(CreateUser(seed.SeedUser))
	fmt.Println(ListUsers.ListUser())
}
