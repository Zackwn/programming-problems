const arr = [1, 1, 0, -1, -1]

const solve = (arr) => {
  let Positive = 0
  let Negative = 0
  let Zero = 0
  arr.forEach((element) => {
    // element > 0 ? Positive++ : element === 0 ? Zero++ : Negative++

    // positive
    if (element > 0) {
      Positive++
      // zero
    } else if (element === 0) {
      Zero++
      // negative
    } else {
      Negative++
    }
  })
  const result = [  
    Positive / arr.length,
    Negative / arr.length,
    Zero / arr.length
  ]
  result.forEach(element => console.log(element))
  // return result
}

console.log(solve(arr))