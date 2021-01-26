'use strict';

const fs = require('fs');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', inputStdin => {
  inputString += inputStdin;
});

process.stdin.on('end', _ => {
  inputString = inputString.replace(/\s*$/, '')
    .split('\n')
    .map(str => str.replace(/\s*$/, ''));

  main();
});

function readLine() {
  return inputString[currentLine++];
}

/*
  Calcula a distância através da ditância de duas space stations dividida por dois, 
  junto a distancia dá primeira cidade e da última.
  A primeira e última cidades não são calculadas juntos com as demais pois não podem
  ser rodeadas por space stations.
  Para a primeira cidade = |firstCity - firstSpaceStation|
  Para a última cidade = |lastCity - lastSpaceStation|

  Algoritimo:
  * coloca as space stations em ordem
  * faz o cálculo das space stations / 2
  * intera os cálculos com a última e primeira cidade 
  * retorna o maior valor dos cálculos
*/
function flatlandSpaceStations(n, c) {
  let distances = []

  c.sort((a, b) => a - b)

  for (let i = 0; i < c.length - 1; i++) {
    const distance = Math.floor(Math.abs(c[i + 1] - c[i]) / 2)
    distances.push(distance)
  }

  const firstCityDistance = Math.abs(c[0] - 0)
  const lastCityDistance = Math.abs(n - c[c.length - 1]) - 1

  return Math.max(...distances, firstCityDistance, lastCityDistance)
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const nm = readLine().split(' ');

  const n = parseInt(nm[0], 10);

  const m = parseInt(nm[1], 10);

  const c = readLine().split(' ').map(cTemp => parseInt(cTemp, 10));

  let result = flatlandSpaceStations(n, c);

  ws.write(result + "\n");

  ws.end();
}