package leet_code

func increasingBST(root *TreeNode) *TreeNode {
	var data []int
	increasingBSTV1(root, &data)
	var d *TreeNode
	var cur *TreeNode
	for _, v := range data {
		if d == nil {
			d = &TreeNode{Val: v}
			cur = d
		} else {
			cur.Right = &TreeNode{Val: v}
			cur = cur.Right
		}
	}
	return d
}

func increasingBSTV1(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	increasingBSTV1(root.Left, data)
	*data = append(*data, root.Val)
	increasingBSTV1(root.Right, data)

}
