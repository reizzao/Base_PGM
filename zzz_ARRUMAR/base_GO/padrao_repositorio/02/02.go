package main

import (
	"fmt"
)

// 1 - Contrato_Principal_NovoTipo: contrato do novotipo
type Entity struct {
	All   []Props
	Props Props
}
type Props struct {
	Texto1 string
	Texto2 string
}

func (i Entity) CreateEntity(d Props) Entity {
	slice := append(i.All, d)
	res := Entity{
		All:   slice,
		Props: d,
	}
	return res
}

// 2 - Instancia_do_Novotipo : instanciar o Contrato_Principal_NovoTipo, aqui pode ser memory, ou externo a entrada das informacoes.
var input1 = Props{
	Texto1: "Texto_Memory_Um",
	Texto2: "Texto_Memory_Dois",
}

var input2 = Props{
	Texto1: "Texto_Json_Um",
	Texto2: "Texto_Json_Dois",
}

// var repo_1 =  CreateEntity(input1)

// var repo_2 = CreateEntity(input2)

// 3 - Metodo: dar acoes ao novotipo - Não precisa fazer metodo para instancias - somente funcoes de repositorio

// 4 - Repositorio - contrato de acoes que o principal pode fazer
// todo: na real nao to usando esta instancia : IRepoPrincipal como obj - mas esta tendo outro uso

type IRepoPrincipal interface {
	Iandar() string
	Ivooar() string
	IComputed() string
}

func (i Entity) Iandar() string {
	return fmt.Sprintf("%s andou", i.Props.Texto1)
}

func (i Entity) Ivooar() string {
	return fmt.Sprintf("%s voou", i.Props.Texto1)
}

func (i Entity) IComputed() string {
	return fmt.Sprintf("Computed 2 campos >> %s %s ", i.Props.Texto1, i.Props.Texto2)
}

// INUSE - Repo - usando um repositorio apartir de uma escolha - que vem da Instancia_do_Novotipo

type Iinuse struct {
	Andar    string
	Voar     string
	Computed string
}

// Espelho do interface Repo - na real está dando chaves a cada metodo do Repo
var inuse_1 = func(e Entity) Iinuse {
	res := Iinuse{
		Andar:    e.Iandar(),
		Voar:     e.Ivooar(),
		Computed: e.IComputed(),
	}

	return res

}

// Criando o controlador de Repo_em_Uso
var constrol_INUSE = repo_2
var constrol_INUSE_2 = repo_1

func main() {

	// Usando o Controlador de Repo_em_Uso -
	fmt.Println(inuse_1(constrol_INUSE).Andar)
	fmt.Println(inuse_1(constrol_INUSE_2).Voar)
	fmt.Println(inuse_1(constrol_INUSE_2).Computed)

}
