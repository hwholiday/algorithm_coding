package leet_code

type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	for head != nil {
		node := head.Next
		head.Next = cur
		cur = head
		head = node
	}
	return cur
}
