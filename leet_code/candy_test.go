package leet_code

import (
	"testing"
)

func candy(ratings []int) (d int) {
	l := len(ratings)
	start := make([]int, l)
	for i := 0; i < l; i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			start[i] = start[i-1] + 1
		} else {
			start[i] = 1
		}
	}
	var end int
	for i := l - 1; i >= 0; i-- {
		if i < l-1 && ratings[i] > ratings[i+1] {
			end++
		} else {
			end = 1
		}
		d += max(end, start[i])
	}
	return d
}

func Test_candy(t *testing.T) {
	t.Log(candy([]int{1, 0, 2}))
}
