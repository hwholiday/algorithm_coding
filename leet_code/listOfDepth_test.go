package leet_code

import (
	"fmt"
	"testing"
)

//面试题 04.03. 特定深度节点链表
func listOfDepth(tree *TreeNode) []*ListNode {
	h := getTreeNodeH(tree)
	var (
		dataListNode = make([]*ListNode, h)
	)
	rangeTreeNode([]*TreeNode{tree}, 0, dataListNode)
	return dataListNode
}

func PrintListNode(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func rangeTreeNode(tree []*TreeNode, j int, listNode []*ListNode) (data []*TreeNode) {
	if tree == nil {
		return nil
	}
	var cur *ListNode
	for i := 0; i < len(tree); i++ {
		if tree[i] == nil {
			continue
		}
		if tree[i].Left != nil {
			data = append(data, tree[i].Left)
		}
		if tree[i].Right != nil {
			data = append(data, tree[i].Right)
		}
		if listNode[j] == nil {
			listNode[j] = &ListNode{Val: tree[i].Val}
			cur = listNode[j]
		} else {
			cur.Next = &ListNode{Val: tree[i].Val}
			cur = cur.Next
		}
	}
	j++
	return rangeTreeNode(data, j, listNode)
}

func getTreeNodeH(tree *TreeNode) int {
	if tree == nil {
		return 0
	}
	return max(getTreeNodeH(tree.Left), getTreeNodeH(tree.Right)) + 1
}

func Test_listOfDepth(t *testing.T) {
	data := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   8,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	listOfDepth(data)

}
