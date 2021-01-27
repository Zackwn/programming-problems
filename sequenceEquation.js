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

function permutationEquation(p) {
  const result = new Array()
  for (let i = 0; i < p.length; i++) {
    result[p[p[i] - 1] - 1] = i + 1
  }
  return result
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const n = parseInt(readLine(), 10);

  const p = readLine().split(' ').map(pTemp => parseInt(pTemp, 10));

  let result = permutationEquation(p);

  ws.write(result.join("\n") + "\n");

  ws.end();
}