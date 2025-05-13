package main

import "fmt"

// 1- Campos Entidade
type Enty struct {
	Id        string
	Nome      string
	SobreNome string
}
type CreateEntyRequestDTO struct {
	Nome      string
	SobreNome string
}

type IEntyRepository interface {
	// All []Enty
	Create(d Enty) Enty
}

// Implementacao Contrato Interface Repo
func (t Enty) Create(d Enty) Enty{
	res := d
	return res
}

// func CreateEntyRepository(d Enty) IEntyRepository{
// 	res := IEntyRepository{
// 		Create: CreateEnty(d),
// 	}
// 	return res
// }

// 3  -
// func CreateEnty(d Enty) Enty {
// 	return d
// }

// 2 - Funcionalidade_Usecase
func CreateEntyUsecase(d CreateEntyRequestDTO,) Enty {
	// regras
	enty := Enty{
		Id:        "makeID()--1",
		Nome:      d.Nome,
		SobreNome: d.SobreNome,
	}

	saved := enty.Create(enty)
	return saved
}

// objetos de uso
// fake - request
var FakeCreateEntyRequestDTO1 = CreateEntyRequestDTO{
	Nome:      "Reinaldo",
	SobreNome: "Zacharias",
}
var FakeCreateEntyRequestDTO2 = CreateEntyRequestDTO{
	Nome:      "Gustavo",
	SobreNome: "Eduardo",
}

// ultimo ProMain
func TesterEnty() {
	fmt.Println(CreateEntyUsecase(FakeCreateEntyRequestDTO1))
}

func main() {
	TesterEnty()
}
