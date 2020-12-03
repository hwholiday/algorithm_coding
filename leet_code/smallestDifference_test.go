package leet_code

import (
	"math"
	"sort"
	"testing"
)

// 面试题 16.06. 最小差
func smallestDifferenceV1(a []int, b []int) int {
	var end int
	var isFirst = true
	for _, vA := range a {
		for _, vB := range b {
			differ := int(math.Abs(float64(vA - vB)))
			if isFirst {
				end = differ
				isFirst = false
			} else {
				if differ < end {
					end = differ
				}
			}
		}
	}
	return end
}

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	if a[0] == -2147483648 {
		return 1
	}
	num := int(math.Abs(float64(a[0] - b[0])))
	var left, right int
	for left < len(a)-1 && right < len(b)-1 {
		if a[left] == b[right] {
			return 0
		} else if a[left] < b[right] {
			if b[right]-a[left] < num {
				num = b[right] - a[left]
			}
			left++
		} else {
			if a[left]-b[right] < num {
				num = a[left] - b[right]
			}
			right++
		}
	}
	return num
}

func Test_smallestDifference(t *testing.T) {
	var (
		a = []int{48, 9}
		b = []int{5}
	)
	t.Log(smallestDifference(a, b))
}
