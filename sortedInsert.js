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

const DoublyLinkedListNode = class {
  constructor(nodeData) {
    this.data = nodeData;
    this.next = null;
    this.prev = null;
  }
};

const DoublyLinkedList = class {
  constructor() {
    this.head = null;
    this.tail = null;
  }

  insertNode(nodeData) {
    let node = new DoublyLinkedListNode(nodeData);

    if (this.head == null) {
      this.head = node;
    } else {
      this.tail.next = node;
      node.prev = this.tail;
    }

    this.tail = node;
  }
};

function printDoublyLinkedList(node, sep, ws) {
  while (node != null) {
    ws.write(String(node.data));

    node = node.next;

    if (node != null) {
      ws.write(sep);
    }
  }
}

// Complete the sortedInsert function below.

/*
 * For your reference:
 *
 * DoublyLinkedListNode {
 *     int data;
 *     DoublyLinkedListNode next;
 *     DoublyLinkedListNode prev;
 * }
 *
 */
function sortedInsert(head, data) {
  let newNode = new DoublyLinkedListNode(data)
  if (!head) {
    head = newNode
    return head
  }
  if (data < head.data) {
    newNode.next = head
    head.prev = newNode
    return newNode
  }

  let currentNode = head

  while (data > currentNode.data) {
    let refNode = currentNode
    currentNode = currentNode.next
    if (refNode.next === null) {
      refNode.next = newNode
      newNode.next = null
      newNode.prev = refNode
      return head
    }
  }

  let leadingPosition = currentNode.prev
  leadingPosition.next = newNode
  newNode.next = currentNode
  currentNode.prev = newNode
  newNode.prev = leadingPosition

  return head
}

function main() {
  const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

  const t = parseInt(readLine(), 10);

  for (let tItr = 0; tItr < t; tItr++) {
    const llistCount = parseInt(readLine(), 10);

    let llist = new DoublyLinkedList();

    for (let i = 0; i < llistCount; i++) {
      const llistItem = parseInt(readLine(), 10);
      llist.insertNode(llistItem);
    }

    const data = parseInt(readLine(), 10);

    let llist1 = sortedInsert(llist.head, data);

    printDoublyLinkedList(llist1, " ", ws)
    ws.write("\n");
  }

  ws.end();
}
