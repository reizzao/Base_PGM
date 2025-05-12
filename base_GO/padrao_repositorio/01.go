package main

import "fmt"

// 1. Interface
type EntyInterface interface {
	Add()
	Delete()
}

// 2. Implementation
type EntyInterfaceImpl struct{} // tipo da implementacao :

// implementacoes Memory
func (i EntyInterfaceImpl) Add() {
	fmt.Println("Memory adicionado Enty")
}

func (i EntyInterfaceImpl) Delete() {
	fmt.Println("Memory deletado Enty")
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
	EntyInterfaceImpl
}

func (i EntyInterfaceImpl_2) Add() {
	fmt.Println("External_1 adicionado Enty")
}

func (i EntyInterfaceImpl_2) Delete() {
	fmt.Println("External_1 deletado Enty")
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

func ProMainEnty() {
	RepoINUSE.Add()
	RepoINUSE.Delete()
}

// Main
func main() {
	ProMainEnty()
}
