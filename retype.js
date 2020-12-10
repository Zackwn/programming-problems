function retype(n, k, s) {
  let takenTime = k + (n - s)

  const backwardsTakenTime = Math.abs(k - s)
  // restartTakenTime = s

  if (backwardsTakenTime > s) {
    takenTime += s
  } else {
    takenTime += backwardsTakenTime
  }

  return takenTime
}

console.log(retype(10, 5, 2))
console.log(retype(10, 7, 6))