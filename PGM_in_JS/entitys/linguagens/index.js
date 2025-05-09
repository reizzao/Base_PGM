

class Types {
  constructor(o) {
    this.conceito = o.conceito;
  }
}

function createTypes(d) {
  const res = new Types(d)
  return res
}

export { createTypes }


