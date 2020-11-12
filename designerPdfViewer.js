function designerPdfViewer(h, word) {
  const alphabet = ["a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"]
  let greaterHeight = 0
  word.split('').forEach(letter => {
    let heightIndex = alphabet.indexOf(letter)
    if (h[heightIndex] > greaterHeight) {
      greaterHeight = h[heightIndex]
    }
  })
  return greaterHeight * word.length
}