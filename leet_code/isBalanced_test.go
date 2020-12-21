package leet_code

// 平衡二叉树
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(TreeNodeHeight(root.Left)-TreeNodeHeight(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func TreeNodeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(TreeNodeHeight(root.Left)+1, TreeNodeHeight(root.Right)+1)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
