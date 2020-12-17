package leet_code

type ListNodeA struct {
	Val  int
	Next *ListNodeA
}

//反转链表
func reverseList(head *ListNodeA) *ListNodeA {
	var cur *ListNodeA
	for head != nil {
		node := head.Next
		head.Next = cur
		cur = head
		head = node
	}
	return cur
}
