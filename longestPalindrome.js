/**
 * @param {string} s
 * @return {string}
 */

var longestPalindrome = function (s) {

  let res = ''
  for (let index = 0; index < s.length; index++) {
    // odd length
    let right = index
    let left = index
    while (left >= 0 && right < s.length && s.charAt(left) === s.charAt(right)) {
      if ((right - left + 1) > res.length) {
        res = s.slice(left, right + 1)
      }
      left--
      right++
    }
    // even length
    right = index + 1
    left = index
    while (left >= 0 && right < s.length && s.charAt(left) === s.charAt(right)) {
      if ((right - left + 1) > res.length) {
        res = s.slice(left, right + 1)
      }
      left--
      right++
    }
  }

  return res
};