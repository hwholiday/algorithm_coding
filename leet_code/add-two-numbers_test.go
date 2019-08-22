package leet_code

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Log(addTwoNumbers(&ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}, &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

//两数相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var data = new(ListNode)
	return data
}
