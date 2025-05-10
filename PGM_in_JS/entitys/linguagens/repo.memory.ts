import type { ITypes } from "./index.ts"
import type { ITypeRepositories } from "./types.ts";

class TypesRepoMemory implements ITypeRepositories {

  public readonly all: ITypes[] = []

  create(d: ITypes): ITypes {
    this.all.push(d)
    return d
  }

  list(): ITypes[] {
    return this.all
  }

}

const typesRepoMemory = new TypesRepoMemory()

export { typesRepoMemory }