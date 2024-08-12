package entity

type Linguagem struct{
	Projeto Projeto
}

type Projeto struct{
	ComandosProjeto ComandosProjeto
}

type ComandosProjeto struct{
	Inicio_Modulo string
	Dependencias string
}