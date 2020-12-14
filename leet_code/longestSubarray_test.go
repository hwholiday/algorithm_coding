package leet_code

import (
	"testing"
)

//删掉一个元素以后全为 1 的最长子数组
func longestSubarray(nums []int) int {
	var (
		sum   = 0
		left  = 0
		right = 0
		max   = 0
	)
	for right = 0; right < len(nums); right++ {
		sum += nums[right]
		for left < right && sum <= right-left-1 {
			if nums[left] == 1 {
				sum--
			}
			left++
		}
		if max < right-left {
			max = right - left
		}
	}
	return max
}

func longestSubarrayV2(nums []int) int {
	var (
		sum   = 0
		left  = 0
		right = 0
	)
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			left++
			right++
			if sum < left {
				sum = left
			}
		} else {
			left = right
			right = 0
		}
	}
	if sum == len(nums) {
		sum--
	}
	return sum
}

func Test_longestSubarray(t *testing.T) {
	var data = []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
	t.Log(longestSubarray(data))
}
