function whatFlavors(cost: Array<number>, money: number) {
  const costMap = new Map()
  for (let index = 0; index < cost.length; index++) {
    const complain = money - cost[index]
    if (costMap.has(complain)) {
      const twoCosts = [costMap.get(complain), index + 1]
        .sort((a, b) => a - b)

      return twoCosts.join(' ')
    }
    costMap.set(cost[index], index + 1)
  }
}

console.log(whatFlavors([1, 4, 5, 3, 2], 4))