package main

// Singly-linked lists are already defined with this interface:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

//
func removeKFromList(l *ListNode, k int) *ListNode {
	if l == nil {
		return l
	}
	for l != nil && l.Value == k {
		l = l.Next
	}
	node := l
	for node != nil && node.Next != nil {
		if node.Next.Value == k {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return l
}
