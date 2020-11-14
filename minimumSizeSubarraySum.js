const arr = [2, 3, 1, 2, 4, 3, 7]

const solve = (intArr, s) => {
  let start = 0
  let sum = 0
  let ans = Number.MAX_VALUE
  console.log(intArr)
  for (let end = 0; end < intArr.length; end++) {
    sum += intArr[end]
    console.log({ plus: sum })
    while (sum >= s) {
      ans = Math.min(ans, end - start + 1)
      sum -= intArr[start++]
      console.log({ minus: sum })
    }
  }
  return ans === Number.MAX_VALUE ? 0 : ans
}

console.log(solve(arr, 7)) // 1