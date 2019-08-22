package leet_code

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{2, 7, 11, 15}, 6))
}

//两数相和
func twoSum(nums []int, target int) []int {
	var data []int
	var endFor = false
	for index, val := range nums {
		for i := index + 1; i < len(nums); i++ {
			if val+nums[i] == target {
				data = append(data, index, i)
				endFor = true
				break
			}
		}
		if endFor {
			break
		}
	}
	return data
}
