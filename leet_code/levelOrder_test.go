package leet_code

import (
	"testing"
)

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var d = make([][]int, getTreeNodeH(root))
	rangeLevelOrder([]*TreeNode{root}, 0, d)
	return d
}

func rangeLevelOrder(root []*TreeNode, j int, d [][]int) (data []*TreeNode) {
	if len(root) <= 0 {
		return
	}
	dInt := make([]int, len(root))
	d[j] = dInt
	for i := 0; i < len(root); i++ {
		if root[i] == nil {
			continue
		}
		if root[i].Left != nil {
			data = append(data, root[i].Left)
		}
		if root[i].Right != nil {
			data = append(data, root[i].Right)
		}
		d[j][i] = root[i].Val
	}
	j++
	return rangeLevelOrder(data, j, d)
}

func Test_levelOrder(t *testing.T) {
	t.Log(levelOrder(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}))
}
