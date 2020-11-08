// [1, 2, 3, 4, 5]
const chaotic = [2, 1, 5, 3, 4]

const minimumBrides = (arr) => {
  let totalBrides = 0

  let expectedFirst = 1
  let expectedSecond = 2
  let expectedThird = 3

  for (let index = 0; index < arr.length; index++) {
    if (arr[index] === expectedFirst) {
      expectedFirst = expectedSecond
      expectedSecond = expectedThird
      expectedThird += 1
    } else if (arr[index] === expectedSecond) {
      totalBrides += 1
      expectedSecond = expectedThird
      expectedThird += 1
    } else if (arr[index] === expectedThird) {
      totalBrides += 2
      expectedThird += 1
    } else {
      console.log('Too chaotic')
    }
  }
  
  console.log(totalBrides)
}

minimumBrides(chaotic)