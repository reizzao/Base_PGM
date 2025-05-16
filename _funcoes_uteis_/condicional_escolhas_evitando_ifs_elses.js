const optionsData = {
  azul: "O escolhido é >> azul.",
  amarelo: "O escolhido é >> amarelo.",
  vermelho: "O escolhido é >> vermelho.",
};

/**
 * @desc acao que da opcoes previstas de respostas atraves de uma escolha
 * @input a escolha desejada pelo consumidor
 * @return um guardado objeto com as opcoes[tendo na posicao - o input recebido] || uma resposta default em caso de erro ou nao previsto
 * @logica :: -em uma funcao por parametro :receba a escolha do consumidor.
- guarde um objeto onde suas props sejam as opcoes que vais dar ao consumidor escolher com seu retorno desejado a oferece, obs: as props tem que ter o nome do que sera disponibilizado.
- devolva esse guardado[com o recebido na posicao] ou uma resposta default caso seja nenhuam das opcoes previtas.'
 */

function CondicionaisPrevistas(d) {
  const res = optionsData;
  return res[d] || "Resposta_Default :: vc escolheu uma opção não prevista!";
}

// uso :
console.log(CondicionaisPrevistas("azul"));
console.log(CondicionaisPrevistas("vermelho"));
console.log(CondicionaisPrevistas("qualquer uma"));
