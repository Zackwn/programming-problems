/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
  for (let index = 0; index < nums.length; index++) {
    const number = nums[index]
    const need = target - number
    const indexOfNeedInNums = nums.indexOf(need, index + 1)
    if (indexOfNeedInNums !== -1) {
      return [index, indexOfNeedInNums]
    }
  }
};