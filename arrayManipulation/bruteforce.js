function arrayManipulation(n, queries) {
  let result = new Array(n).fill(0)
  queries.forEach(query => {
    for (let index = query[0]; index < query[1] + 1; index++) {
      result[index - 1] += query[2]
    }
  })
  return Math.max(...result)
}