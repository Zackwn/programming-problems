function solve(s, e) {
  let boringNumbers = []
  for (let index = s; index <= e; index++) {
    let isBoringNumber = true
    index.toString().split('').map(Number).forEach((n, index) => {
      // even position
      if ((index + 1) % 2 === 0) {
        if (n % 2 !== 0) {
          isBoringNumber = false
        }
        // odd position
      } else {
        if (n % 2 === 0) {
          isBoringNumber = false
        }
      }
    })
    if (isBoringNumber) {
      boringNumbers.push(index)
    }
  }
  return boringNumbers
}

console.log(solve(5, 15).length)
console.log(solve(120, 125).length)
console.log(solve(779, 783).length)