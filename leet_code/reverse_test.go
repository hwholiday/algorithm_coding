package leet_code

import (
	"math"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Log(reverse(-123))
}
func reverse(x int) int {
	var (
		min = math.Pow(-2, 31)
		max = math.Pow(2, 31) - 1
	)
	var i float64
	for x != 0 {
		i = i*10 + float64(x%10)
		x = x / 10
	}
	if min > i || max < i {
		return 0
	}
	return int(i)
}
