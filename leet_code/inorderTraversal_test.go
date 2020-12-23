package leet_code

//二叉树的中序遍历
func inorderTraversal(root *TreeNode) (data []int) {
	inorderTraversalV2(root, &data)
	return
}

func inorderTraversalV2(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	inorderTraversalV2(root.Left, data)
	*data = append(*data, root.Val)
	inorderTraversalV2(root.Right, data)
}
