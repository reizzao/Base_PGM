package seed

import (
	"github.com/pgm/api/entity"
)

var Universal_Projeto_Seed = entity.Projetos_UNV{
	ConvencaoPastas: entity.ConvencaoPastas{
		NomePastaEntidade: 1,
		Nome_Arquivo:      1,
	},
	OrdemFluxoEntidade: []entity.OrdemFluxoEntidade{
		entity.OrdemFluxoEntidade{

			Contrato_Formato_Entidade: entity.OrdemFluxoEntidadeProps{
				Ordem:      1,
				Finalidade: Seed_Sinificado_Contrato_Formato_Entidade,
			},

			Usecase_AcaoCasoDeUso: entity.OrdemFluxoEntidadeProps{
				Ordem:      2,
				Finalidade: Seed_Significado_USECASE,
			},

			Seed: entity.OrdemFluxoEntidadeProps{
				Ordem:      3,
				Finalidade: Seed_Significado_SEED,
			},

			Tester: entity.OrdemFluxoEntidadeProps{
				Ordem:      4,
				Finalidade: Seed_Significado_TESTER,
			},

			Test: entity.OrdemFluxoEntidadeProps{
				Ordem:      5,
				Finalidade: Seed_Significado_TEST,
			},

			Index_Modulo: entity.OrdemFluxoEntidadeProps{
				Ordem:      6,
				Finalidade: Seed_Significado_INDEX_MODULO,
			},

			Main_Root: entity.OrdemFluxoEntidadeProps{
				Ordem:      7,
				Finalidade: Seed_Significado_MAIN_ROOT,
			},
		},
	},
}
