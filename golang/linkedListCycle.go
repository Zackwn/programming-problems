/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	node := head
	specialnum := 100001
	for node != nil && node.Next != nil {
		if node.Val == specialnum {
			return true
		}
		node.Val = specialnum
		node = node.Next
	}
	return false
}
