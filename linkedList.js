var node = /** @class */ (function () {
    function node(data) {
        this.data = data;
    }
    return node;
}());
var List = /** @class */ (function () {
    function List() {
        this.length = 0;
    }
    List.prototype.getCurrentAndPreviusNodes = function () {
        var currentNode = this.head;
        var previusNode = null;
        while (currentNode.next) {
            previusNode = currentNode;
            currentNode = currentNode.next;
        }
        return {
            current: currentNode,
            previus: previusNode
        };
    };
    List.prototype.add = function (value) {
        if (!this.head) {
            this.head = new node(value);
            this.length++;
            return;
        }
        var current = this.getCurrentAndPreviusNodes().current;
        current.next = new node(value);
        this.length++;
    };
    Object.defineProperty(List.prototype, "last", {
        get: function () {
            return this.getCurrentAndPreviusNodes().current.data;
        },
        enumerable: false,
        configurable: true
    });
    Object.defineProperty(List.prototype, "first", {
        get: function () {
            return this.head.data;
        },
        enumerable: false,
        configurable: true
    });
    List.prototype.pop = function () {
        var _a = this.getCurrentAndPreviusNodes(), previus = _a.previus, current = _a.current;
        var data = current.data;
        delete previus.next;
        this.length--;
        return data;
    };
    List.prototype.removeDuplicates = function () {
        var dataHistory = [];
        var currentNode = this.head;
        var previusNode = currentNode;
        while (currentNode) {
            if (dataHistory.includes(currentNode.data)) {
                // there is a duplicated value, so remove currentNode
                previusNode.next = currentNode.next;
                currentNode = previusNode.next;
                continue;
            }
            dataHistory.push(currentNode.data);
            previusNode = currentNode;
            currentNode = currentNode.next;
        }
        console.log(this.head);
    };
    List.prototype.forEach = function (callbackFN) {
        var current = this.head;
        while (current) {
            callbackFN(current.data);
            current = current.next;
        }
    };
    List.prototype.map = function (callbackFN) {
        var current = this.head;
        var result = [];
        while (current) {
            var returnData = callbackFN(current.data);
            if (returnData !== undefined) {
                result.push(returnData);
            }
            current = current.next;
        }
        return result;
    };
    List.prototype.reduce = function (callbackFN) {
        var current = this.head;
        var accumulator = current.data;
        current = current.next;
        while (current) {
            accumulator = callbackFN(accumulator, current.data);
            current = current.next;
        }
        return accumulator;
    };
    return List;
}());
var stringLogs = false;
var myStringList = new List();
myStringList.add('test1');
myStringList.add('test2');
myStringList.add('test3');
myStringList.forEach(function (v) { return stringLogs && console.log(v); });
var newArray = myStringList.map(function (v) {
    return v + " newArray";
});
myStringList.pop();
stringLogs && console.log(newArray);
stringLogs && console.log(myStringList.first);
stringLogs && console.log(myStringList.last);
stringLogs && console.log(myStringList.length);
var numberLogs = false;
var myNumberList = new List();
myNumberList.add(1);
myNumberList.add(2);
myNumberList.add(2);
myNumberList.add(3);
console.log(myNumberList.head);
myNumberList.removeDuplicates();
// const sum = myNumberList.reduce((acc, v) => acc + v)
// numberLogs && console.log(sum)
