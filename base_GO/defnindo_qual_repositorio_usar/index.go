package main

import (
	"fmt"
)

// 1 - Contrato_Principal_NovoTipo: contrato do novotipo
type Principal struct {
	Texto1 string
	Texto2 string
}

// 2 - Instancia_do_Novotipo : instanciar o Contrato_Principal_NovoTipo, aqui pode ser memory, ou externo a entrada das informacoes.
var i_principal_memory = Principal{
	Texto1: "Texto_Memory_Um",
	Texto2: "Texto_Memory_Dois",
}

var i_principal_json = Principal{
	Texto1: "Texto_Json_Um",
	Texto2: "Texto_Json_Dois",
}

// 3 - Metodo: dar acoes ao novotipo - Não precisa fazer metodo para instancias - somente funcoes de repositorio

// 4 - Repositorio - contrato de acoes que o principal pode fazer
type IRepoPrincipal interface {
	Iandar() string
}

func (i Principal) Iandar() string {
	return fmt.Sprintf("%s andou", i.Texto1)
}
func (i Principal) Ivooar() string {
	return fmt.Sprintf("%s voou", i.Texto1)
}

// INUSE - Repo - usando um repositorio apartir de uma escolha - que vem da Instancia_do_Novotipo

type Iinuse struct {
	Andar string
	Voar  string
}

var inuse_1 = func(principal Principal) Iinuse {
	res := Iinuse{
		Andar: principal.Iandar(),
		Voar:  principal.Ivooar(),
	}

	return res

}

// Criando o controlador de Repo_em_Uso
// var constrol_INUSE = i_principal_memory
var constrol_INUSE = i_principal_json

func main() {

	// Usando o Controlador de Repo_em_Uso -
	fmt.Println(inuse_1(constrol_INUSE).Andar)
	fmt.Println(inuse_1(constrol_INUSE).Voar)

}
