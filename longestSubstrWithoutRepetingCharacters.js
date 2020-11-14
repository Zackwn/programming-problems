let string = 'abcad'

const solve = (str) => {
  if (!str) return 0
  let max = 1
  let start = 0
  const substrSet = new Set()
  for (let end = 0; end < str.length; end++) {
    const currentLetter = str.charAt(end)
    while (substrSet.has(currentLetter)) {
      substrSet.delete(str.charAt(start++))
    }
    substrSet.add(currentLetter)
    max = Math.max(max, end - start + 1)
  }
  console.log(substrSet)
  return max
}

console.log(solve(string))

// export { }