import { ITypes } from "./index.ts";

interface ITypeRepositories {
  all: ITypes[]
  create(d: ITypes): ITypes
  list(): ITypes[]
}

export type { ITypeRepositories }