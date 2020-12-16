package leet_code

import "testing"

//面试题 04.12. 求和路径
func pathSum(root *TreeNode, sum int) int {
	var end = new(int)
	pathNode(root, sum, end)
	return *end
}
func pathNode(root *TreeNode, sum int, end *int) {
	if root == nil {
		return
	}
	dfs(root, sum, end)
	pathNode(root.Left, sum, end)
	pathNode(root.Right, sum, end)
}

func dfs(root *TreeNode, sum int, end *int) {
	if root == nil {
		return
	}
	if root.Val == sum {
		*end++
	}
	dfs(root.Left, sum-root.Val, end)
	dfs(root.Right, sum-root.Val, end)
}

func Test_pathSum(t *testing.T) {

}
