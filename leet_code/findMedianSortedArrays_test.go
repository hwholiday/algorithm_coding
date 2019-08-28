package leet_code

import (
	"fmt"
	"sort"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Log(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

//寻找两个有序数组的中位数
//中位数，就是数组排序后处于数组最中间的那个元素。
// 数数有几个数，加 1，再除以 2
// 如果数组长度是奇数，最中间就是位置为（n+1）／2的那个元素。
// 如果是偶数呢，标准的定义是位置为n/2和位置为n/2+1的两个元素的和除以2的结果，有些复杂。为了解答的方便，我们假设数组长度暂且都为奇数吧。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	len := len(nums1) + 1
	var v1 float64
	if (len-1)%2 == 0 {
		v1 = (float64(nums1[len/2-1]) + float64(nums1[len/2+1-1])) / 2
	} else {
		v1 = float64(nums1[(len+1)/2-1])
	}
	fmt.Println(v1)
	return v1
}
