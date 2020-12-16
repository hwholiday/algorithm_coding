package leet_code

import (
	"fmt"
	"testing"
)

//两个链表的第一个公共节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapListNode := make(map[ListNode]*ListNode)
	for headA != nil {
		mapListNode[*headA] = headA
		headA = headA.Next
	}
	for headB != nil {
		if mapListNode[*headB] == headB {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

func Test_getIntersectionNode(t *testing.T) {
	data := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	datab := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 8,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
	}
	d := getIntersectionNode(data, datab)
	if d != nil {
		fmt.Print(d.Val)
		t.Log(d.Val)
	}
}
