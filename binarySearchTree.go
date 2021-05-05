package main

import (
	"fmt"
)

func main() {
	bst := BinarySearchTree{}

	// bst.Insert(2)
	// bst.Insert(3)
	// bst.Insert(1)
	// bst.Insert(9)
	// bst.Insert(8)

	bst.Insert(30)
	bst.Insert(20)
	bst.Insert(25)
	bst.Insert(26)
	bst.Insert(21)

	// bst.Insert(30)
	bst.Delete(25)

	fmt.Println(bst.head.left.right.Value) // 26
	fmt.Println(bst.Search(26))            // true
	fmt.Println(bst.Search(25))            // false
}

type Node struct {
	Value int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	head *Node
}

func (b *BinarySearchTree) Delete(value int) {
	var previusNode *Node
	node := b.head
	for node != nil {
		if node.Value == value {
			if node.right != nil {
				l := node.right
				for l.left != nil {
					l = l.left
				}
				l.left = node.left
			}
			previusNode.right = node.right
			return
		}
		previusNode = node
		if value >= node.Value {
			node = node.right
		} else {
			node = node.left
		}
	}
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
