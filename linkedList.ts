class node<T> {
  data: T
  next: node<T>
  constructor(data: T) {
    this.data = data
  }
}

class List<T> {
  private head: node<T>
  length: number

  constructor() {
    this.length = 0
  }

  private getCurrentAndPreviusNodes(): {current: node<T>, previus: node<T>} {
    let currentNode = this.head
    let previusNode = null
    while(currentNode.next) {
      previusNode = currentNode
      currentNode = currentNode.next
    }
    return {
      current: currentNode,
      previus: previusNode
    }
  }

  add(value: T) {
    if (!this.head) {
      this.head = new node(value)
      this.length++
      return
    }
    let { current } = this.getCurrentAndPreviusNodes()
    current.next = new node(value)
    this.length++
  }

  get last(): T | undefined {
    return this.getCurrentAndPreviusNodes().current.data
  }

  get first(): T | undefined {
    return this.head.data
  }

  pop() {
    let { previus, current } = this.getCurrentAndPreviusNodes()
    let data = current.data
    delete previus.next
    this.length--
    return data
  }

  forEach(callbackFN: (value: T) => any) {
    let current = this.head
    while(current) {
      callbackFN(current.data)
      current = current.next
    }
  }

  map(callbackFN: (value: T) => T | T[] | any) {
    let current = this.head
    const result = []
    while(current) {
      let returnData = callbackFN(current.data)
      if (returnData !== undefined) {
        result.push(returnData)
      }
      current = current.next
    }
    return result
  }

  reduce(callbackFN: (accumulator: T, currentValue: T) => T) {
    let current = this.head
    let accumulator: T | T[] = current.data
    current = current.next
    while(current) {
      accumulator = callbackFN(accumulator, current.data)
      current = current.next
    }
    return accumulator
  }
}

let stringLogs: boolean = false 
const myStringList = new List<string>()
myStringList.add('test1')
myStringList.add('test2')
myStringList.add('test3')
myStringList.forEach((v) => stringLogs && console.log(v))
const newArray = myStringList.map((v) => {
  return `${v} newArray`
})
myStringList.pop()
stringLogs && console.log(newArray)
stringLogs && console.log(myStringList.first)
stringLogs && console.log(myStringList.last)
stringLogs && console.log(myStringList.length)

let numberLogs: boolean = true
const myNumberList = new List<number>()
myNumberList.add(1)
myNumberList.add(2)
myNumberList.add(3)
const sum = myNumberList.reduce((acc, v) => acc + v)
numberLogs && console.log(sum)