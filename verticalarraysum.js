const arr = [
  [11, 2, 4],
  [4, 5, 6],
  [10, 8, -12]
]

const sumArray = (arr) => {
  return arr.reduce((accumulator, currentValue) => {
    return accumulator + currentValue
  })
}

const solve = (arr) => {
  const DiagonalOneArr = []
  const DiagonalTwoArr = []
  let indexOne = 0
  let indexTwo = arr.length - 1
  for (let index = 0; index < arr.length; index++) {
    DiagonalOneArr.push(arr[index][indexOne])
    DiagonalTwoArr.push(arr[index][indexTwo])
    indexOne += 1
    indexTwo -= 1
  }
  return Math.abs(sumArray(DiagonalOneArr) - sumArray(DiagonalTwoArr))
}

console.log(solve(arr))