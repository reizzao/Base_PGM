package seed

import "github.com/pgm/api/entity"

var Seed_Significado_SEED = entity.Significado{
	Nome:           "Seed",
	Significado:    "Informacoes Default por padrao de fabrica d eum objeto",
	Aplicabilidade: []string{"Usa-se para preencher por default informacoes de OBJETOS."},
}

var Seed_Sinificado_Contrato_Formato_Entidade = entity.Significado{
	Nome:           "Contrato Formato Entidade",
	Significado:    "È o Contrato de como será o formato dos campos do Objeto.",
	Aplicabilidade: []string{"struct", "class"},
}

var Seed_Significado_USECASE = entity.Significado{
	Nome:           "UseCase - Casos de Uso",
	Significado:    "São acoes que podemoms fazer com a entidade objeto",
	Aplicabilidade: []string{"CRUD", "create", ""},
}

var Seed_Significado_TEST = entity.Significado{
	Nome:           "Test",
	Significado:    "Testar a Entidade",
	Aplicabilidade: []string{"Subir a app sem erros", "garantir que mudanças na estrutura nao afetem outras partes da app ja testados"},
}

var Seed_Significado_INDEX_MODULO = entity.Significado{
	Nome:           "Index_Modulo",
	Significado:    "Concentrar um indice do modulo de funcoes casos de uso(usecases) por entidade especifica",
	Aplicabilidade: []string{""},
}

var Seed_Significado_MAIN_ROOT = entity.Significado{
	Nome:           "Main_Root",
	Significado:    "Concentrar um indice de todos Index_Modulo",
	Aplicabilidade: []string{""},
}
