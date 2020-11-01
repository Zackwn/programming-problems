"use strict";
// import { inspect } from 'util'
exports.__esModule = true;
var Node = /** @class */ (function () {
    function Node(data) {
        this.data = data;
        this.left = null;
        this.right = null;
    }
    return Node;
}());
/**
  Every BinarySearchTree have one root,
  from it the greater numbers will be on the right side and the smallers on the left.

  Exemplo:

                      right "12"
          right "9"
                      left "1"
  root "8"
                      right "5"
          left "4"
                      left "2"


  As seen above, even one element not being directly the root (begin),
  he can origin two more forks witch follow the same root's guide
 */
var BinarySearchTree = /** @class */ (function () {
    function BinarySearchTree() {
    }
    BinarySearchTree.prototype.add = function () {
        var data = [];
        for (var _i = 0; _i < arguments.length; _i++) {
            data[_i] = arguments[_i];
        }
        var root = null;
        if (this.lastRoot) {
            root = this.lastRoot;
        }
        for (var index = 0; index < data.length; index++) {
            root = this.insert(root, data[index]);
        }
        this.lastRoot = root;
    };
    BinarySearchTree.prototype.insert = function (root, data) {
        if (!root) {
            this.root = new Node(data);
            return this.root;
        }
        // smaller or equal -> left
        if (data <= root.data) {
            // if the left fork already exists, continue from it
            if (root.left) {
                return this.insert(root.left, data);
            }
            else {
                root.left = new Node(data);
            }
            // maior -> direita
        }
        else {
            // if the right fork alredy exits, continue from it
            if (root.right) {
                return this.insert(root.right, data);
            }
            else {
                root.right = new Node(data);
            }
        }
        return this.root;
    };
    Object.defineProperty(BinarySearchTree.prototype, "getHeight", {
        get: function () {
            return this.calculateHeight(this.root);
        },
        enumerable: false,
        configurable: true
    });
    /**
      The height of a binarySearchTree is the distance between the root and its furthest leaf.
    */
    BinarySearchTree.prototype.calculateHeight = function (root) {
        var rightCount = 0;
        var leftCount = 0;
        if (!root) {
            return 0;
        }
        if (root.left) {
            leftCount = this.calculateHeight(root.left) + 1;
        }
        if (root.right) {
            rightCount = this.calculateHeight(root.right) + 1;
        }
        return Math.max(rightCount, leftCount);
    };
    return BinarySearchTree;
}());
/* Testing */
var tree = new BinarySearchTree();
var data = [3, 5, 2];
var data2 = [1, 4, 6, 7];
tree.add.apply(tree, data);
tree.add.apply(tree, data2);
// console.log(inspect(tree.root, { showHidden: false, depth: null }))
console.log(tree.getHeight); // -> 3
