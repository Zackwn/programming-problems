const solve = (a, b) => {
  /*
    Factors of X -> numbers that X is divisible by.
    Is a factor of X -> Is divisible by X
  */
  let count = 0
  for (let x = 1; x <= 100; x++) {
    // check if any (x) is divisible by any (a) element
    if (a.every(number => (x % number === 0))) {
      // check if any (b) element is divisible by (x) witch can be divisible by all (a) elements
      if (b.every(number => (number % x === 0))) {
        count++
      }
    }
  }
  return count
}

console.log(solve([2, 4], [16, 32, 96]))