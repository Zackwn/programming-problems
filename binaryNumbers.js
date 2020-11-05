"use strict";
exports.__esModule = true;
var Solution = /** @class */ (function () {
    function Solution(number) {
        this.decimalNumber = number;
        this.binaryNumber = this.convertDecimalToBinary();
    }
    Solution.prototype.convertDecimalToBinary = function () {
        var binary = '';
        for (var decimalNumber = this.decimalNumber; decimalNumber > 0; decimalNumber = (decimalNumber - (decimalNumber % 2)) / 2) {
            var rest = decimalNumber % 2;
            binary = "" + rest + binary;
        }
        return Number(binary);
    };
    Solution.prototype.calculateConsecutiveOnes = function () {
        var binaryStr = this.binaryNumber.toString();
        var lettersMap = binaryStr.match(/([0-1])\1*/g) || [];
        var consecutiveOnes = 0;
        lettersMap.forEach(function (str) {
            if (str.includes('1')) {
                if (str.length > consecutiveOnes) {
                    consecutiveOnes = str.length;
                }
            }
        });
        return consecutiveOnes;
    };
    Object.defineProperty(Solution.prototype, "consecutiveOnes", {
        get: function () {
            return this.calculateConsecutiveOnes();
        },
        enumerable: false,
        configurable: true
    });
    return Solution;
}());
var solve = new Solution(5);
console.log(solve.consecutiveOnes);
