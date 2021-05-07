export function xo(str: string) {
  // your code here
  let xcount = 0
  let ocount = 0
  for (let i = 0; i < str.length; i++) {
    if (str[i].toLowerCase() === "x") {
      xcount++
    }
    if (str[i].toLowerCase() === "o") {
      ocount++
    }
  } 
  return xcount == ocount
}