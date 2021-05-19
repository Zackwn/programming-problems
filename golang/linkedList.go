package main

import "fmt"

func (this *MyLinkedList) Display() {
	fmt.Println("")

	fmt.Println(this.head)
	n := this.head
	for n != nil {
		fmt.Print(n.value, " ")
		n = n.next
	}
	fmt.Println("")
	fmt.Println("")
}

type Node struct {
	next  *Node
	value int
}

type MyLinkedList struct {
	head   *Node
	length int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}

	cur := this.head
	for index > 0 {
		index--
		cur = cur.next
	}
	if cur == nil {
		return -1
	}

	return cur.value
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.head == nil {
		this.head = &Node{
			value: val,
		}
		this.length++
		return
	}

	n := Node{
		next:  this.head,
		value: val,
	}
	this.head = &n
	this.length++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.head == nil {
		this.head = &Node{
			value: val,
		}
		this.length++
		return
	}

	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = &Node{
		value: val,
	}

	this.length++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.length {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.length {
		this.AddAtTail(val)
		return
	}

	cur := this.head
	previus := cur
	for index > 0 {
		previus = cur
		cur = cur.next
		index--
	}

	n := &Node{
		next:  cur,
		value: val,
	}
	previus.next = n

	this.length++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.length {
		return
	}

	cur := this.head
	var previus *Node
	for index > 0 {
		previus = cur
		cur = cur.next
		index--
	}

	if previus == nil {
		this.head = cur.next
	} else {
		previus.next = cur.next
	}

	this.length--
}

/**
* Your MyLinkedList object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.Get(index);
* obj.AddAtHead(val);
* obj.AddAtTail(val);
* obj.AddAtIndex(index,val);
* obj.DeleteAtIndex(index);
 */
