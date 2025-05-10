package CreateUser

import (
	"githubendereco/Dto"
	"githubendereco/controlators"
	t "githubendereco/entitys"
)

func PrepareCreateUser(u Dto.CreateUserDTO) t.User {
	form := t.User{
		Id:        "uuid()--1",
		Nome:      u.Nome,
		Sobrenome: u.Sobrenome,
	}

	res := controlators.UserRepository.Save(form)

	return res
}
