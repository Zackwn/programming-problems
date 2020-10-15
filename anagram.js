"use strict";
var query = ['abc', 'lock', 'qaui'];
var dictionary = ['cba', 'abc', 'ckol', 'kolc', 'kcol', 'lolck', 'qauil'];
var sort = function (target) {
    // return target.split('').sort().join('')
    var targetArr = target.split('').map(function (letter) { return letter.toUpperCase(); });
    var alphabet = ["A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"];
    var result = '';
    for (var ALPindex = 0; ALPindex < alphabet.length; ALPindex++) {
        var letter = alphabet[ALPindex];
        var indexesInTarget = [];
        for (var i = 0; i < targetArr.length; i++) {
            if (targetArr[i] === letter) {
                indexesInTarget.push(i);
            }
        }
        if (indexesInTarget.length === 0) {
            continue;
        }
        indexesInTarget.forEach(function (index) {
            result += target[index];
        });
    }
    return result;
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
