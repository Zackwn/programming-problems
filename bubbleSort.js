function countSwaps(a) {
  let swaps = 0
  for (let index = 0; index < a.length; index++) {
    for (let aheadIndex = index + 1; aheadIndex < a.length; aheadIndex++) {
      if (a[index] > a[aheadIndex]) {
        let temp = a[index]
        a[index] = a[aheadIndex]
        a[aheadIndex] = temp
        swaps++
      }
    }

  }
  console.log(`Array is sorted in ${swaps} swaps.`)
  console.log(`First Element: ${a[0]}\nLast Element: ${a[a.length - 1]}`)
}

countSwaps([1, 2, 3]) // 0 swaps