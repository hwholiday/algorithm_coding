package leet_code

import (
	"sort"
	"testing"
)

//三数之和
func TestThreeSum(t *testing.T) {
	t.Log(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	var (
		result [][]int
	)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					data := []int{nums[i], nums[j], nums[k]}
					sort.Ints(data)
					for _, v := range result {
						if !equal(v, data) {
							result = append(result, data)
						}
					}
				}
			}
		}
	}
	return result
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
