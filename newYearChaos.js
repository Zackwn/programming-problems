// [1, 2, 3, 4, 5]
const chaotic = [2, 1, 5, 3, 4]

const minumumBribes = (arr) => {
  let totalBribes = 0

  let expectedFirst = 1
  let expectedSecond = 2
  let expectedThird = 3

  for (let index = 0; index < arr.length; index++) {
    if (arr[index] === expectedFirst) {
      expectedFirst = expectedSecond
      expectedSecond = expectedThird
      expectedThird += 1
    } else if (arr[index] === expectedSecond) {
      totalBribes += 1
      expectedSecond = expectedThird
      expectedThird += 1
    } else if (arr[index] === expectedThird) {
      totalBribes += 2
      expectedThird += 1
    } else {
      console.log('Too chaotic')
    }
  }
  
  console.log(totalBribes)
}

minumumBribes(chaotic)