function minimumSwapsAscending(arr: number[]) {
  let minSwaps = 0

  for (let i = 0; i < arr.length; i++) {

    // the expected number
    const right = i + 1

    // if the current element is no the expected number
    if (arr[i] !== right) {

      // get where (index) the expected number is
      const rightIdx = arr.indexOf(right, i)

      // swap the expected to the current
      arr[rightIdx] = arr[i]

      // swap the current to the spected 
      arr[i] = right

      // increse one swap
      ++minSwaps
    }
  }
  return minSwaps
}

function minimumSwapsDesceding(arr: number[]) {
  let minimumSwaps = 0
  let greaterElement = Math.max(...arr)
  for (let index = arr.length - 1; index > 0; index--) {
    let right = Math.abs(greaterElement - index)
    if (arr[index] !== right) {
      let rightIndex = arr.indexOf(right)
      let tempEl = arr[index]
      arr[index] = right
      arr[rightIndex] = tempEl
      minimumSwaps++
    }
  }
  return minimumSwaps
}

minimumSwapsAscending([4, 3, 1, 2]) // 3
minimumSwapsDesceding([1, 2, 3, 4]) // 2

export { }