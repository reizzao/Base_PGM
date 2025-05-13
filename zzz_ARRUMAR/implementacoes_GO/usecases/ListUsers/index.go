package ListUsers

import (
	"githubendereco/controlators"
	"githubendereco/entitys"
)

func ListUser() []*entitys.User {
	res := controlators.UserRepositoryINUSE.List()
	return res
}

// todo esta retornando somente o array vazio
