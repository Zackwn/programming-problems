/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
  const mergedArray = [...nums1, ...nums2].sort((a, b) => a - b)
  if (mergedArray.length % 2 === 0) {
    // even median
    const midLeft = mergedArray[Math.floor((mergedArray.length - 1) / 2)]
    const midRight = mergedArray[Math.floor((mergedArray.length) / 2)]
    return (midLeft + midRight) / 2
  }
  // odd median
  return mergedArray[Math.floor((mergedArray.length - 1) / 2)]
};