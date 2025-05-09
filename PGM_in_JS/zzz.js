class NumerosRepoInMemory {
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

const instRepoInUse = new NumerosRepoInMemory();

class Controller {
  delRepo = instRepoInUse;

  create(o) {
    this.delRepo.create(o);
  }

  list() {
    this.delRepo.list();
  }

  useSomaTotal_UltimosInseridos() { // todo fail
    this.delRepo.useSomaTotal_UltimosInseridos();
  }
}

const newController = new Controller();

const fakeObjNumeros01 = { n1: 1000, n2: 2000 };
const fakeObjNumeros02 = { n1: 3000, n2: 4000 };

newController.create(fakeObjNumeros01);
newController.create(fakeObjNumeros02);

// todo :mandando o Controle delegar
// console.log(newController.list());
// console.log(newController.ultimoInsertArray());
// console.log(instRepoInUse.useSomaTotal_UltimosInseridos()); // todo fail

// ok fazendo pela instancia do repo
console.log(instRepoInUse.list());
console.log(instRepoInUse.ultimoInsertArray());
console.log(instRepoInUse.useSomaTotal_UltimosInseridos()); // todo fail
