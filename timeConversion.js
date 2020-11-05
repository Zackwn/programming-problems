// const hourToConvert = '12:01:00AM'
// const hourToConvert = '12:01:00PM' 
const hourToConvert = '07:05:45PM'
// const hourToConvert = '12:40:22AM'
// const hourToConvert = '06:40:03AM'
// const hourToConvert = '04:59:59AM'
// const hourToConvert = '12:45:54PM'

const solve = (hourToConvert) => {
  const AM_PM = hourToConvert.slice(-2)
  const Time = hourToConvert.slice(0, -2).split(':')

  if (AM_PM === 'AM' && Time[0] === '12') {
    Time[0] = '00'
    return Time.join(':')
  }

  if (AM_PM === 'PM') {
    Time[0] = `${(Number(Time[0]) % 12) + 12}`
  }

  return Time.join(':')
}

console.log(solve(hourToConvert))