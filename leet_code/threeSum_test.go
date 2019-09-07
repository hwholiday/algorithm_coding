package leet_code

import (
	"fmt"
	"sort"
	"testing"
)

//三数之和
func TestThreeSum(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	var (
		result    [][]int
		numsMap   = make(map[int]int, len(nums))
		resultMap = make(map[string]string)
		allIsZero = true
	)
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums)
	for i, v := range nums {
		numsMap[v] = i
		if v != 0 {
			allIsZero = false
		}
	}
	if allIsZero {
		result = append(result, []int{0, 0, 0})
		return result
	}

	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			if index, ok := numsMap[0-(v+nums[j])]; ok {
				data := []int{v, nums[j], 0 - (v + nums[j])}
				sort.Ints(data)
				if _, ok := resultMap[fmt.Sprintf("%d_%d_%d", data[0], data[1], data[2])]; !ok {
					if index != i && index != j {
						result = append(result, data)
					}
					resultMap[fmt.Sprintf("%d_%d_%d", data[0], data[1], data[2])] = ""
				}
			}
		}
	}
	return result
}
