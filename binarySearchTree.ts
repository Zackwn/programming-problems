// import { inspect } from 'util'

class Node<T> {
  data: T
  left: Node<T>
  right: Node<T>

  constructor(data: T) {
    this.data = data
    this.left = null
    this.right = null
  }
}

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
class BinarySearchTree<T = number> {
  root: Node<T>
  private lastRoot: Node<T>

  constructor() { }

  add(...data: Array<T>): void {
    let root = null
    if (this.lastRoot) {
      root = this.lastRoot
    }
    for (let index = 0; index < data.length; index++) {
      root = this.insert(root, data[index])
    }
    this.lastRoot = root
  }

  private insert(root: Node<T>, data: T) {
    if (!root) {
      this.root = new Node(data)
      return this.root
    }

    // smaller or equal -> left
    if (data <= root.data) {
      // if the left fork already exists, continue from it
      if (root.left) {
        return this.insert(root.left, data)
      } else {
        root.left = new Node(data)
      }
      // maior -> direita
    } else {
      // if the right fork alredy exits, continue from it
      if (root.right) {
        return this.insert(root.right, data)
      } else {
        root.right = new Node(data)
      }
    }

    return this.root
  }

  get getHeight() {
    return this.calculateHeight(this.root)
  }

  /**
    The height of a binarySearchTree is the distance between the root and its furthest leaf.
  */
  private calculateHeight(root: Node<T>) {
    let rightCount = 0
    let leftCount = 0

    if (!root) {
      return 0
    }

    if (root.left) {
      leftCount = this.calculateHeight(root.left) + 1
    }
    if (root.right) {
      rightCount = this.calculateHeight(root.right) + 1
    }

    return Math.max(rightCount, leftCount)
  }
}

/* Testing */
const tree = new BinarySearchTree()

let data = [3, 5, 2]
let data2 = [1, 4, 6, 7]

tree.add(...data)
tree.add(...data2)

// console.log(inspect(tree.root, { showHidden: false, depth: null }))
console.log(tree.getHeight) // -> 3

export { }