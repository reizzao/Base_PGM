package main

import "fmt"

// 1. Interface
type EntyInterface interface {
	Add(p DTO)
	Delete()
	NomeCompleto() string
}

type CollectionEnty = []DTO

type DTO struct {
	Nome      string
	SobreNome string
}

// 2. Implementation
type EntyInterfaceImpl struct {
	Props DTO
	CollectionEnty
}

// tipo da implementacao :

// implementacoes Memory
func (i EntyInterfaceImpl) Add(d DTO) {

	_slice := append(i.CollectionEnty, d)
	_ = _slice
	fmt.Println("Memory adicionado Enty", d.Nome)
}

func (i EntyInterfaceImpl) Delete() {
	fmt.Println("Memory deletado Enty")
}

func (i EntyInterfaceImpl) NomeCompleto() string {
	return fmt.Sprintf("Memory o NomeCompleto é: %s %s", i.Props.Nome, i.Props.SobreNome)
}

// 3. Construct

func EntyRepositoryMemory() EntyInterfaceImpl {
	return EntyInterfaceImpl{}
}

var ObjEntyImportantRepository = EntyRepositoryMemory()

// -- Outro Arquivo
// implementacoes External_1

// tipo da implementacao :
type EntyInterfaceImpl_2 struct {
	// aqui cabe as props da entidade
	Props DTO
	EntyInterfaceImpl
}

func (i EntyInterfaceImpl_2) Add(p DTO) {
	fmt.Println("External_1 adicionado Enty", p.Nome)
}

func (i EntyInterfaceImpl_2) Delete() {
	fmt.Println("External_1 deletado Enty")
}

func (i EntyInterfaceImpl_2) NomeCompleto() string {
	return fmt.Sprintf("External o NomeCompleto é: %s %s", i.EntyInterfaceImpl.Props.Nome, i.EntyInterfaceImpl.Props.SobreNome)
}

func EntyRepositoryexternal_1() EntyInterfaceImpl_2 {
	return EntyInterfaceImpl_2{}
}

var ObjEntyRepositoryexternal_1 = EntyRepositoryexternal_1()

// Fim -- implementacoes External_1 ---

// -- Outro Arquivo -- Fim

// 4. Important - Control Repository :: by useMain

// Important : Escolha do Repo em USO

type Option struct {
	Memory     EntyInterfaceImpl
	External_1 EntyInterfaceImpl_2
}

var OptionsEntyRepository = Option{
	Memory:     ObjEntyImportantRepository,
	External_1: ObjEntyRepositoryexternal_1,
}

var RepoINUSE = OptionsEntyRepository.Memory

// var RepoINUSE = OptionsEntyRepository.External_1

var dto_1 = DTO{
	Nome:      "Reinaldo",
	SobreNome: "Zacharias",
}

func ProMainEnty() {
	RepoINUSE.Add(dto_1)
	RepoINUSE.Delete()
	fmt.Println(RepoINUSE.NomeCompleto())
}

// Main
func main() {
	ProMainEnty()
}
