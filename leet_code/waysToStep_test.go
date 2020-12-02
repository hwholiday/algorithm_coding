package leet_code

import (
	"testing"
)

func waysToStep(n int) int {
	var d = make([]int, n+4)
	d[0] = 0
	d[1] = 1
	d[2] = 2
	d[3] = 4
	for i := 4; i <= n; i++ {
		d[i] = d[i-1] + d[i-2] + d[i-3]
		d[i] = d[i] % 1000000007
	}
	return d[n]
}

func Test_waysToStep(t *testing.T) {
	t.Log(waysToStep(5))
}
