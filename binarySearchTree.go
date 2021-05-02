package main

import "fmt"

func main() {
	bst := BinarySearchTree{}

	bst.Insert(2)
	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(9)
	bst.Insert(8)

	fmt.Println(bst.Search(8)) // true
	fmt.Println(bst.Search(5)) // false
}

type Node struct {
	Value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	head *Node
}

func (b *BinarySearchTree) Search(value int) bool {
	node := b.head
	for node != nil {
		if node.Value == value {
			return true
		}
		if value >= node.Value {
			node = node.right
		} else {
			node = node.left
		}
	}
	return false
}

func (b *BinarySearchTree) Insert(value int) {
	if b.head == nil {
		b.head = &Node{
			Value: value,
		}
		return
	}

	node := b.head
	for node != nil {
		if value >= node.Value {
			// right
			if node.right != nil {
				node = node.right
			} else {
				node.right = &Node{
					Value: value,
				}
				break
			}
		} else {
			// left
			if node.left != nil {
				node = node.left
			} else {
				node.left = &Node{
					Value: value,
				}
				break
			}
		}
	}
}
