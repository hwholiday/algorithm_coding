package leet_code

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	t.Log(removeNthFromEnd(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}, 2))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n <= 0 || head == nil {
		return head
	}
	fast := head
	//获取要截断后的一个
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	slow := head
	for fast.Next != nil {
		slow = slow.Next
		fmt.Println("slow", slow.Val)
		fast = fast.Next
		fmt.Println("fast", fast.Val)
	}
	slow.Next = slow.Next.Next
	return head
}
