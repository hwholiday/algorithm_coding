package leet_code

import "testing"

// 三角形个数
func triangleNumber(nums []int) int {
	var count int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j] > nums[k] && nums[i]+nums[k] > nums[j] && nums[j]+nums[k] > nums[i] {
					count++
				}
			}
		}
	}
	return count
}

func Test_triangleNumber(t *testing.T) {
	t.Log(triangleNumber([]int{2, 2, 3, 4}))
}
