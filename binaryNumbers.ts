class Solution {
  private decimalNumber: number
  private binaryNumber: number

  constructor(number: number) {
    this.decimalNumber = number
    this.binaryNumber = this.convertDecimalToBinary()
  }

  private convertDecimalToBinary(): number {
    let binary = ''
    for (
      let decimalNumber = this.decimalNumber;
      decimalNumber > 0;
      decimalNumber = (decimalNumber - (decimalNumber % 2)) / 2
    ) {
      const rest = decimalNumber % 2
      binary = `${rest}${binary}`
    }
    return Number(binary)
  }

  private calculateConsecutiveOnes(): any {
    let binaryStr = this.binaryNumber.toString()
    let lettersMap = binaryStr.match(/([0-1])\1*/g)||[]
    let consecutiveOnes = 0
    lettersMap.forEach((str) => {
      if (str.includes('1')) {
        if (str.length > consecutiveOnes) {
          consecutiveOnes = str.length
        }
      }
    })
    return consecutiveOnes
  }

  get consecutiveOnes(): number {
    return this.calculateConsecutiveOnes()
  }
}

const solve = new Solution(5)
console.log(solve.consecutiveOnes)

export {}