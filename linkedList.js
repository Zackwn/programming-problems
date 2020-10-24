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
    return List;
}());
var myList = new List();
myList.add('test1');
myList.add('test2');
myList.add('test3');
myList.pop();
console.log(myList.first); // test1
console.log(myList.last); // test2
console.log(myList.length); // 2
