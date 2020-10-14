"use strict";
var query = ['abc', 'lock', 'qaui'];
var dictionary = ['cba', 'abc', 'ckol', 'kolc', 'kcol', 'lolck', 'qauil'];
var sort = function (target) {
    return target.split('').sort().join('');
};
var solve = function (query, dictionary) {
    var result = {};
    query.forEach(function (queryElement) {
        var queryElementSorted = sort(queryElement);
        result[queryElement] = 0;
        dictionary.forEach(function (dictionaryElement) {
            var dictonaryElementSorted = sort(dictionaryElement);
            if (queryElementSorted === dictonaryElementSorted) {
                result[queryElement] += 1;
            }
        });
    });
    var response = [];
    for (var key in result) {
        response.push(result[key]);
    }
    return response;
};
console.log(solve(query, dictionary));
