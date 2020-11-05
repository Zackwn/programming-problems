let arr = [1, 2, 3, 4, 5]
// let arr = [5, 5, 5, 5, 5]

const sumArrayToNumber = (arr, ...numbers) => {
  return arr.concat(numbers).reduce((accumulator, currentElement) => {
    return accumulator + currentElement
  })
}

const solve = (arr) => {
  let greaterNumber = 0
  let smallestNumber = arr[0]
  for (let index = 0; index < arr.length; index++) {
    const currentElement = arr[index]
    if (currentElement > greaterNumber) {
      greaterNumber = currentElement
    }
    if (currentElement < smallestNumber) {
      smallestNumber = currentElement
    }
  }
  arr.splice(arr.indexOf(greaterNumber), 1)
  arr.splice(arr.indexOf(smallestNumber), 1)
  const max = sumArrayToNumber(arr, greaterNumber)
  const min = sumArrayToNumber(arr, smallestNumber)
  console.log(min, max)
  // return null
}

console.log(solve(arr))