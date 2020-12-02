package leet_code

import (
	"sort"
	"testing"
)

//1296. 划分数组为连续数字的集合
func isPossibleDivide(nums []int, k int) bool {
	if k == 1 {
		return true
	}
	sort.Ints(nums)
	var mapNums = make(map[int]int)
	for _, v := range nums {
		mapNums[v]++
	}
	for _, v := range nums {
		val := mapNums[v]
		if val == 0 {
			continue
		}
		mapNums[v]--
		for i := 1; i < k; i++ {
			key := v + i
			va := mapNums[key]
			if va == 0 {
				return false
			}
			mapNums[key]--
		}
	}
	return true
}

func TestIspossibledivide(t *testing.T) {
	var nums = []int{1, 2, 3, 3, 4, 4, 5, 6}
	t.Log(isPossibleDivide(nums, 4))
}
