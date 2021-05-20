package main

// Singly-linked lists are already defined with this interface:
type ListNode struct {
	Value interface{}
	Next  *ListNode
}

//
func isListPalindrome(l *ListNode) bool {
	values := make([]interface{}, 0)
	node := l
	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}
	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-1-i] {
			return false
		}
	}
	return true
}
