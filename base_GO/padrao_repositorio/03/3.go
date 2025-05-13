package main

type Enty struct{
	Nome string
	SobreNome string
}

type EntyRepository struct{
	All []Enty
	Create func(d Enty) Enty
}



func main() {}