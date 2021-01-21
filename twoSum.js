function whatFlavors(cost, money) {
    var costMap = new Map();
    for (var index = 0; index < cost.length; index++) {
        var complain = money - cost[index];
        if (costMap.has(complain)) {
            var twoCosts = [costMap.get(complain), index + 1]
                .sort(function (a, b) { return a - b; });
            return twoCosts.join(' ');
        }
        costMap.set(cost[index], index + 1);
    }
}
console.log(whatFlavors([1, 4, 5, 3, 2], 4));
