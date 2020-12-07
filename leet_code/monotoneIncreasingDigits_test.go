package leet_code

import (
	"strconv"
	"strings"
	"testing"
)

//单调递增的数字
func monotoneIncreasingDigits(N int) int {
	var (
		s = strconv.Itoa(N)
	)
	var result []string
	var i = 0
	for i = 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			break
		}
	}
	if i == len(s)-1 {
		return N
	}
	var addN bool
	for _, v := range s {
		if addN {
			result = append(result, "9")
		} else {
			if v == int32(s[i]) {
				addN = true
				k, _ := strconv.Atoi(string(s[i]))
				result = append(result, strconv.Itoa(k-1))
			} else {
				result = append(result, string(v))
			}
		}
	}

	k, _ := strconv.Atoi(strings.Join(result, ""))
	return k
}

func Test_monotoneIncreasingDigits(t *testing.T) {
	t.Log(monotoneIncreasingDigits(668841))
}
