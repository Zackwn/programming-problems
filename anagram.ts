const query = ['abc', 'lock', 'qaui']
const dictionary = ['cba', 'abc', 'ckol', 'kolc', 'kcol', 'lolck', 'qauil']

const sort = (target: string) => {
  // return target.split('').sort().join('')
  const targetArr = target.split('').map(letter => letter.toUpperCase())
  const alphabet = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"]
  let result = ''
  for (let ALPindex = 0; ALPindex < alphabet.length; ALPindex++) {
    const letter = alphabet[ALPindex]
    const indexesInTarget = []
    for (let i = 0; i < targetArr.length; i++) {
      if (targetArr[i] === letter) {
        indexesInTarget.push(i)
      }
    }
    if (indexesInTarget.length === 0) {
      continue
    }
    indexesInTarget.forEach(index => {
      result += target[index]
    })
  }
  return result
}

interface ResultInterface {
  [key: string]: number
}

const solve = (query: string[], dictionary: string[]) => {
  let result: ResultInterface = {}
  query.forEach(queryElement => {
    const queryElementSorted = sort(queryElement)
    result[queryElement] = 0
    dictionary.forEach(dictionaryElement => {
      const dictonaryElementSorted = sort(dictionaryElement)
      if (queryElementSorted === dictonaryElementSorted) {
        result[queryElement] += 1
      }
    })
  })
  const response: number[] = []
  for (let key in result) {
    response.push(result[key])
  }
  return response
}

console.log(solve(query, dictionary))