package seed

import "github.com/pgm/api/entity"

var Linguagem_Golang_Seed = entity.Linguagem{
	Projeto: entity.Projeto{
		ComandosProjeto: entity.ComandosProjeto{
			Inicio_Modulo: "`go mod init github.com/nomeprojeto`",
			Dependencias:  "`go mod tidy`",
		},
	},
}
