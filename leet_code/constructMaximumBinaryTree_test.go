package leet_code

import (
	"testing"
)

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return maxTreeNode(nums, 0, len(nums)-1)
}

func maxTreeNode(num []int, i, j int) *TreeNode {
	if i > j {
		return nil
	}
	modI := findMaxIndex(num, i, j)
	root := new(TreeNode)
	root.Val = num[modI]
	root.Left = maxTreeNode(num, i, modI-1)
	root.Right = maxTreeNode(num, modI+1, j)
	return root
}

func findMaxIndex(num []int, i, j int) int {
	var (
		index    = -1
		indexNum = -1
	)
	for k := i; k <= j; k++ {
		if indexNum < num[k] {
			indexNum = num[k]
			index = k
		}
	}
	return index
}

func Test_constructMaximumBinaryTree(t *testing.T) {
	num := []int{3, 2, 1, 6, 0, 5}
	constructMaximumBinaryTree(num)
}
