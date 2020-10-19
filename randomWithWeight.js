"use strict";
class Solution {
    constructor(intArr, logs = false) {
        this.IndexWeight = new Map();
        this.logs = logs;
        this.intArr = intArr;
        this.Indexes = [];
        this.init();
    }
    init() {
        // weight is based on the integer divided by the arr sum
        const intArrSum = this.intArr.reduce((acc, value) => acc + value);
        this.intArr.forEach((int, index) => {
            const porcent = (int / intArrSum) * 100;
            this.IndexWeight.set(index, Math.floor((porcent * intArrSum) / 100));
            this.Indexes.push(index);
        });
    }
    get getTotalTimes() {
        let total = 0;
        for (let element of this.IndexWeight) {
            total += element[1];
        }
        return total;
    }
    get getRandom() {
        const randomIndex = this.Indexes[Math.floor(Math.random() * this.Indexes.length)];
        this.logs && console.log({ IndexWeight: this.IndexWeight });
        this.logs && console.log({ randomIndex });
        const randomIndexTimes = this.IndexWeight.get(randomIndex);
        if (!randomIndexTimes) {
            this.IndexWeight.delete(randomIndex);
            return 9.99999;
        }
        if (randomIndexTimes <= 1) {
            this.IndexWeight.delete(randomIndex);
            this.Indexes.splice(randomIndex, 1);
        }
        else {
            this.IndexWeight.set(randomIndex, randomIndexTimes - 1);
        }
        return this.intArr[randomIndex];
    }
}
const solution = new Solution([5, 10], false);
console.log(solution.getTotalTimes);
for (let key = solution.getTotalTimes; key !== 0; key--) {
    console.log('value ' + solution.getRandom);
}
