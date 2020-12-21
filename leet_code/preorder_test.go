package leet_code

import (
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (data []int) {
	preorderV2(root, &data)
	return
}

func preorderV2(root *Node, data *[]int) {
	if root == nil {
		return
	}
	*data = append(*data, root.Val)
	for i := 0; i < len(root.Children); i++ {
		preorderV2(root.Children[i], data)
	}
	return
}

//589. N叉树的前序遍历
func Test_preorder(t *testing.T) {
	t.Log(preorder(&Node{
		Val: 1,
		Children: []*Node{{
			Val: 3,
			Children: []*Node{
				{Val: 5},
				{Val: 6},
			},
		},
			{
				Val: 2,
			},
			{
				Val: 4,
			},
		},
	}))
}
