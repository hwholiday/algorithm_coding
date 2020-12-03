package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(data int) *TreeNode {
	return &TreeNode{
		Val: data,
	}
}

func (n *TreeNode) Insert(newNode *TreeNode) {
	if n.Val == newNode.Val {
		return
	}
	if newNode.Val > n.Val {
		if n.Right == nil {
			n.Right = newNode
		} else {
			n.Right.Insert(newNode)
		}
	} else {
		if n.Left == nil {
			n.Left = newNode
		} else {
			n.Left.Insert(newNode)
		}
	}
}

func (n *TreeNode) Search(data int) *TreeNode {
	if n == nil {
		return nil
	}
	if n.Val == data {
		return n
	}
	if data > n.Val {
		n.Right.Search(data)
	} else {
		n.Left.Search(data)
	}
	return nil
}
