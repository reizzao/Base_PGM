/*
Recurso : Condicionais : DAR AO USER A OPÇÃO DE DAR UM INPUT - QUE ESTEJA ENTRE AS OPCOES DISPONIVEIS EM UM OBJETO > SEM IFS
*/
const objOptions = {
  UM() {
    return "UM";
  },
  DOIS() {
    return "DOIS";
  },
};

function getOption(objOpcoes, posicao) {
  return objOpcoes[posicao]();
}

// AO USAR A FUNCAO E ENVIAR O INPUT - E ESSE INPUT FOR A CHAVE DE UM DOS METODOS DISPONIVEIS O USER TERA O RESULTADO ENTRE OS PRE ESTABELECIDOS.
console.log(getOption(objOptions, "UM"));
