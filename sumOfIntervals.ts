export function sumOfIntervals(intervals: [number, number][]) {
  const values = new Set()
  for (let interval of intervals) {
    for (let i = interval[0]; i < interval[1]; i++) [
      values.add(i)
    ]
  }
  return values.size
}