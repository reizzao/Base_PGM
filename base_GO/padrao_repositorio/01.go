package main

import "fmt"

// 1. Interface
type EntyInterface interface {
	addEnty()
	deleteEnty()
	modifyEnty()
}

// 2. Implementation
type EntyInterfaceImpl struct{}

func (i EntyInterfaceImpl) addEnty() {
	fmt.Println("adicionado Enty")
}

func (i EntyInterfaceImpl) deleteEnty() {
	fmt.Println("deletado Enty")
}

func (i EntyInterfaceImpl) modifyEnty() {
	fmt.Println("modificado Enty")
}

// 3. Construct

func EntyRepositoryMemory() EntyInterfaceImpl {
	return EntyInterfaceImpl{}
}

// 4. Important - Control Repository
var entyImportantRepository = EntyRepositoryMemory()

// 5. useMain

func ProMainEnty() {
	entyImportantRepository.addEnty()
	entyImportantRepository.deleteEnty()
	entyImportantRepository.modifyEnty()
}

// Main
func main() {
	ProMainEnty()
}
