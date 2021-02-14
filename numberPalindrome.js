/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
  const X = x
  let reverseNumber = 0
  while (x > 0) {
    reverseNumber = (reverseNumber * 10) + x % 10
    x = Math.trunc(x / 10)
  }
  return reverseNumber === X
};