package leet_code

import (
	"sort"
	"testing"
)

//数组拆分 I
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var sum int
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

func Test_arrayPairSum(t *testing.T) {
	t.Log(arrayPairSum([]int{1, 4, 3, 2}))
}
