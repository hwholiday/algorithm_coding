package leet_code

import (
	"testing"
)

//二叉搜索树的第k大节点
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var data []int
	kthLargestV1(root, &data)
	qSort(data)
	return data[k-1]
}

func qSort(num []int) {
	if len(num) <= 1 {
		return
	}
	var (
		start = 0
		end   = len(num) - 1
		mod   = num[0]
	)
	for start < end {
		if mod < num[start+1] {
			num[start], num[start+1] = num[start+1], num[start]
			start++
		} else {
			num[start+1], num[end] = num[end], num[start+1]
			end--
		}
	}
	qSort(num[:start])
	qSort(num[start+1:])
}

func kthLargestV1(root *TreeNode, d *[]int) {
	if root == nil {
		return
	}
	*d = append(*d, root.Val)
	kthLargestV1(root.Left, d)
	kthLargestV1(root.Right, d)
}

func Test_kthLargest(t *testing.T) {
}
