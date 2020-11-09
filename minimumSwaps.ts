function minimumSwaps(arr: number[]) {
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

minimumSwaps([4, 3, 1, 2])

export { }