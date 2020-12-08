package leet_code

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

//输出二叉树
func printTree(root *TreeNode) [][]string {
	h := int(getRootHeight(root))
	fmt.Println(h)
	allNodeNum := power(h) - 1
	var data = make([][]string, h)
	for i := 0; i < h; i++ {
		data[i] = make([]string, int(allNodeNum))
	}
	fill(root, &data, 0, 0, int(allNodeNum))
	return data
}
func fill(root *TreeNode, ans *[][]string, h, l, r int) {
	if root == nil {
		return
	}
	mid := (l + r) / 2
	(*ans)[h][mid] = strconv.Itoa(root.Val)
	fill(root.Left, ans, h+1, l, mid-1)
	fill(root.Right, ans, h+1, mid+1, r)
}
func power(i int) float64 {
	if i == 0 {
		return 1
	}
	return 2 * power(i-1)
}

func getRootHeight(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	return 1 + math.Max(getRootHeight(root.Left), getRootHeight(root.Right))
}

func Test_printTree(t *testing.T) {
	t.Log(printTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}))
}
