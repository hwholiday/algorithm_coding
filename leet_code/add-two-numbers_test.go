package leet_code

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	info := addTwoNumbers(&ListNode{
		Val:  5,
		Next: nil,
	}, &ListNode{
		Val:  5,
		Next: nil,
	})
	for info != nil {
		t.Log(info.Val)
		info = info.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var data = new(ListNode)
	var curr = data
	var less int
	for l1 != nil || l2 != nil {
		var (
			x   int = 0
			y   int = 0
			sum     = 0
		)
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}
		sum = less + x + y
		less = sum / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next

	}
	if less > 0 {
		curr.Next = &ListNode{
			Val:  less,
			Next: nil,
		}
	}
	return data.Next
}
