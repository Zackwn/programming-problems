const memo = new Map()
var fibonacci = function (n) {
  if (n == 0 || n == 1)
    return n;
  let r = memo.get(n)
  if (r !== undefined) {
    return r
  }
  r = fibonacci(n - 1) + fibonacci(n - 2);
  memo.set(n, r)
  return r
}