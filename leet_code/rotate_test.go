package leet_code

import (
	"testing"
)

//189. 旋转数组
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		last := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = last
	}
}

func Test_rotate(t *testing.T) {
	var data = []int{-1, -100, 3, 99}
	rotate(data, 2)
	t.Log(data)
}
