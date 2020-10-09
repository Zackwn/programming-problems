let staircaseSize = 6

const generateRepeatString = (word, size) => {
  let str = ''
  for (let index = 0; index < size; index++) {
    str += word
  }
  return str
}

const solve = (size) => {
  for (let key = 1; key <= size; key++) {
    const baseStr = generateRepeatString(' ',size - key)
    console.log(baseStr+generateRepeatString('#', key))
  }
  // return null
}

console.log(solve(staircaseSize))