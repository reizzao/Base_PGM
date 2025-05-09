class Entidade {
  constructor() {
    this.delRepo = new EntidadeRepositorieMemory();
  }

  create(o) {
    return this.delRepo.create(o);
  }

  list() {
    return this.delRepo.list();
  }

  ultimoInsertArray() {
    return this.delRepo.recebidosAll[this.delRepo.recebidosAll.length - 1];
  }

  useSomaTotal_UltimosInseridos() {
    const o = this.ultimoInsertArray();
    const res = this.delRepo.useSomar(o.n1, o.n2);
    return res;
  }
}

const newEntidade = new Entidade();

//

class EntidadeRepositorieMemory {
  recebidosAll = [];

  create(o) {
    this.n1 = o.n1;
    this.n2 = o.n2;

    this.insertRecebido(o);
  }

  insertRecebido(o) {
    this.recebidosAll.push(o);
  }

  list() {
    return this.recebidosAll;
  }

  ultimoInsertArray() {
    return this.recebidosAll[this.recebidosAll.length - 1];
  }

  useSomar(n1, n2) {
    return n1 + n2;
  }

  useSomaTotal_UltimosInseridos() {
    const o = this.ultimoInsertArray();
    const res = this.useSomar(o.n1, o.n2);
    return res;
  }
}

// Client usando a instancia via controller

// const newEntidadeRepositorie = new NumerosRepoInMemory();

const fakeObjNumeros01 = { n1: 1000, n2: 2000 };
const fakeObjNumeros02 = { n1: 3000, n2: 4000 };

newEntidade.create(fakeObjNumeros01);
newEntidade.create(fakeObjNumeros02);

// -- View

// mandando o Controle delegar
console.log(newEntidade.list());
console.log(newEntidade.ultimoInsertArray());
console.log(newEntidade.useSomaTotal_UltimosInseridos()); // todo fail

// fazendo pela instancia do repo
// console.log(newEntidadeRepositorie.list());
// console.log(newEntidadeRepositorie.ultimoInsertArray());
// console.log(newEntidadeRepositorie.useSomaTotal_UltimosInseridos()); // todo fail
