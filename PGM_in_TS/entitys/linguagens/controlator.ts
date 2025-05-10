import type { ITypeRepositories } from "./types.ts"
import type { ISelectorRepositories } from "../../global/types.ts"
import { typesRepoMemory } from "./repo.memory.ts";

const typeRepos: ISelectorRepositories<ITypeRepositories> = {
  inuse: typesRepoMemory,
}

// Important
// selecao do repositorio desta entidade a estar ativo !

const typeRepoinUse = typeRepos.inuse

export { typeRepoinUse }