package leet_code

import (
	"testing"
)

func TestMaxArea(t *testing.T) {
	t.Log(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	var (
		startIndex int
		endIndex   int = len(height) - 1
		max        int
	)
	for i := 0; i < len(height); i++ {
		var (
			min int
		)
		indexStart := startIndex
		indexEnd := endIndex
		if height[startIndex] < height[endIndex] {
			min = height[startIndex]
			startIndex++
		} else {
			min = height[endIndex]
			endIndex--
		}
		if min*(indexEnd-indexStart) >= max {
			max = min * (indexEnd - indexStart)
		}
	}
	return max
}
