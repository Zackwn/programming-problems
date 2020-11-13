function arrayManipulation(n, queries) {
  const params = []
  queries.forEach((query) => {
    const start = query[0]
    const end = query[1]
    const value = query[2]
    params.push({
      key: start,
      value: value
    })
    params.push({
      key: end,
      value: -value
    })
  })
  params.sort((item1, item2) => {
    if (item1.key === item2.key) {
      // item2.value is the end, so it always will be greater that item1.value
      return item2.value - item1.value
    }

    return item1.key - item2.key
  })
  let max = 0
  let sum = 0
  for (let param of params) {
    sum += param.value
    if (sum > max) {
      max = sum
    }
  }
  return max
}