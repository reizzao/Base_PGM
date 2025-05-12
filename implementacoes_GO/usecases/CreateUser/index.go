package CreateUser

import (
	"githubendereco/Dto"
	t "githubendereco/entitys"
)

func CreateUser(u Dto.CreateUserDTO) t.User {
	res := PrepareCreateUser(u)
	return res
}

