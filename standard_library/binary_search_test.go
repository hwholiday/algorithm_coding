package standard_library

import (
	"testing"
)

func BinarySearch(data []int, f int) int {
	var (
		left  = 0
		right = len(data) - 1
	)
	for left <= right {
		mod := (right + left) / 2
		if data[mod] > f {
			right = mod - 1
		} else if data[mod] < f {
			left = mod + 1
		} else {
			return mod
		}
	}
	return -1
}

func QSort(d []int) {
	if len(d) <= 1 {
		return
	}
	var (
		left  = 0
		right = len(d) - 1
	)
	start := d[0]
	for left < right {
		if start > d[left+1] {
			d[left], d[left+1] = d[left+1], d[left]
			left++
		} else {
			d[left+1], d[right] = d[right], d[left+1]
			right--
		}
	}
	QSort(d[:left])
	QSort(d[left+1:])
}

func Test_BinarySearch(t *testing.T) {
	var a = []int{3, 2, 1, 9, 7, 4, 7, 6}
	QSort(a)
	t.Log(a)
	t.Log(BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8}, 3))
}
