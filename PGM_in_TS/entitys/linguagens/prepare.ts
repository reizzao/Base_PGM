import { typeRepoinUse } from "./controlator.ts";
import type { ITypes } from "./index.ts"
import type { ITypeRepositories } from "./types.ts";

class TypesPrepare {

  // public readonly all: ITypes[] = []
  constructor(public readonly repositorie: ITypeRepositories) { }

  create(d: ITypes): ITypes {
    const created = this.repositorie.create(d)
    return created
  }

  list(): ITypes[] {
    return this.repositorie.list()
  }

}

const typesPrepare = new TypesPrepare(typeRepoinUse)

export { typesPrepare }