package entitys

import "fmt"

type User struct {
	Id        string
	Nome      string
	Sobrenome string
}

func (u User) NomeCompleto() {
	fmt.Sprintf("s% s%", u.Nome, u.Sobrenome)
	return
}

