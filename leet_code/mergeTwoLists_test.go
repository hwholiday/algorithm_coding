package leet_code

import "testing"

//TODO 合并两个有序链表
func TestmergeTwoLists(t *testing.T) {
	t.Log(mergeTwoLists(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}, &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}))
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}
