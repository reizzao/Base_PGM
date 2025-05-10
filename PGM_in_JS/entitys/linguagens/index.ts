import { typesPrepare } from "./prepare.ts";


interface ITypes {
  conceito: string
  propriedade_opcional: string
}

class Types {

  private constructor(public readonly props: ITypes,) {
    this.props = props;
  }

  static create(d: ITypes): ITypes {
    const action = typesPrepare.create(d)
    return action
  }

  list(): ITypes[] {
    const res = typesPrepare.list()
    return res
  }

}

const typesFactory = Types

export { typesFactory }
export type { ITypes, Types }


