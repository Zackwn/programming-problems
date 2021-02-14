/**
 * @param {number} x
 * @return {number}
 */
const MAX_INT = Math.pow(2, 31)
var reverse = function (x) {
  let reverse = 0
  while (x !== 0) {
    const lastDigit = x % 10
    x = Math.trunc(x / 10)
    reverse = (reverse * 10) + lastDigit
    if (reverse < (MAX_INT * -1) || reverse > MAX_INT) {
      return 0;
    }
  }
  return reverse
};