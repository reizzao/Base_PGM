let controle_constructor = false;

class Entidade2 {
  constructor(o) {
    if (!controle_constructor) {
      throw new Error("Construtor é privado - instancie pelo metodo statico.");
    }
    controle_constructor = false; // reseta a variavel

    // dando valor aos campos
    this.campo1 = o.campo1;
    this.campo2 = o.campo2;
  }

  static create(o) {
    controle_constructor = true
    return new Entidade2(o);
  }
}

const fakeInput01 = { campo1: 10, campo2: 20 };

const instUm = Entidade2.create(fakeInput01);

// console.log(new Entidade2(fakeInput)); // o erro funcionou
console.log(instUm);
