
//  - Interno: O Cli precisa usar algum recurso e pra isso criamos uma função - Aqui escolho os recursos que posso disponibilizar ao Cli.
class RecursoCli  {
  i = new DepRecursoCliClass()

 Re_OpOla (d) {
  const dep = inuseRepositoyGetRequestRecurso.getArgumentRecurso(d)
 return this.i.OpOla(dep)
 }

 Re_OpComoVai (d) {
  const dep = inuseRepositoyGetRequestRecurso.getArgumentRecurso(d)
 return this.i.OpComoVai(dep)
 }

}

// 1 - Externo: Recebemos o Objeto de Pedido do Cli
class FormatRequest  {

 getDep1(d) {
  return d.dep1
 }
}

class RepositoyGetRequestRecursoMemory {

  getArgumentRecurso (d) {
    return d
  }
}

const inuseRepositoyGetRequestRecurso = new RepositoyGetRequestRecursoMemory()

//  - interno : Onde seria uma função pode ser um método de uma classe , evitaria hell de funções e abriria opções de operações.
class DepRecursoCliClass {
 OpOla (d) {
   return `Ola ${d}`
 }

 OpComoVai (d) {
   return `Como vai ${d}`
 }

 EmStandyBy1 (d) {
   return `EmStandyBy1 ${d}`
 }

}

//  - interno : Passamos um obj que instância um método de recurso ao cli.
const obj_Insta_RecursoCliOptions = new RecursoCli()

const instaFormatRequestRecurso = new FormatRequest()


//  - Interno : Recurso Ofertado.
console.log(
 obj_Insta_RecursoCliOptions.Re_OpOla (
  instaFormatRequestRecurso.getDep1(
    {
      dep1 : "Rei2"
    }
  ))
);
