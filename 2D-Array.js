function hourglassSum(arr) {
  let maxY = 3
  let maxX = 3
  let total = -Infinity
  
  for (let y = 0; y <= maxY; y++) {
      for (let x = 0; x <= maxX; x++) {
          // top sum
          let sum = arr[y][x] + arr[y][x+1] + arr[y][x+2]
          // middle sum
          sum += arr[y+1][x+1]
          // bottom sum
          sum += arr[y+2][x] + arr[y+2][x+1] + arr[y+2][x+2]

          if (sum >  total) {
              total = sum
          }
      }
  }
  return total
}

const arr = [ [ 1, 1, 1, 0, 0, 0 ],
[ 0, 1, 0, 0, 0, 0 ],
[ 1, 1, 1, 0, 0, 0 ],
[ 0, 0, 2, 4, 4, 0 ],
[ 0, 0, 0, 2, 0, 0 ],
[ 0, 0, 1, 2, 4, 0 ] ]

console.log(hourglassSum(arr)) // 19