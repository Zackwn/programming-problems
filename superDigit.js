// 1
function superDigit(n, k) {
  if (n <= 9) {
    return n
  }
  let nSum = 0
  while (n > 0) {
    nSum += n % 10
    n = Math.floor(n / 10)
  }
  return superDigit(nSum * k, 1)
}

// 2
const superDigit = (s, k) => {
  if (s.length == 1) {
    return s
  }
  let sum = 0
  for (let i = 0; i < s.length; i++) {
    sum += parseFloat(k * parseInt(s.charAt(i)))
  }
  return superDigit(sum.toString(), 1)
}