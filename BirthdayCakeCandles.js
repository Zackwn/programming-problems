const candles = [3, 2, 1, 3, 4]

const solve = (candles) => {
  let greaterCandle = 0
  candles.forEach(candle => {
    if (candle > greaterCandle) {
      greaterCandle = candle
    }
  })
  let duplicate = []
  for (let index = 0; index < candles.length; index++) {
    if (greaterCandle === candles[index]) {
      duplicate.push(greaterCandle)
    }
  }
  if (duplicate.length > 1) {
    // console.log(`Candle height ${greaterCandle}`)
    return duplicate.length
  }
  const nextCandles = candles.filter((candle) => {
    return candle !== greaterCandle
  })
  return solve(nextCandles)
}

console.log(solve(candles))