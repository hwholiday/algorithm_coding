package tree

type RadixTreeNode struct {
	path     string
	children []*RadixTreeNode
}

func NewRadixTreeNode() RadixTreeNode {
	return RadixTreeNode{
		path:     "",
		children: nil,
	}
}
