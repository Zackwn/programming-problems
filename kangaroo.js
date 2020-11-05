const solve = (x1, v1, x2, v2) => {
  let canCatchUp = x1 > x2 ? v1 < v2 : v1 > v2
  let response = 'NO'
  while(canCatchUp) {
    const distance = Math.abs(x1 - x2)
    if (distance === 0) {
      response = 'YES'
      break
    }
    x1 += v1
    x2 += v2
    canCatchUp = x1 > x2 ? v1 < v2 : v1 > v2
  }
  return response
  // if (v1 > v2) {
  //   const remainder = (x1 - x2) % (v1 - v2)
  //   if (remainder === 0) {
  //     return 'YES'
  //   }
  // }
  // return 'NO'
}

// solve(0, 3, 4, 2)
// solve(0, 2, 5, 3)
// solve(21, 6, 47, 3)
console.log(solve(70, 2, 43, 2))