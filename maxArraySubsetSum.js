function maxSubsetSum(arr) {
  let prev1 = 0
  let prev2 = 0
  for (let value of arr) {
      let temp = prev1
      prev1 = Math.max(prev1, value + prev2)
      prev2 = temp
  }
  return prev1
}
