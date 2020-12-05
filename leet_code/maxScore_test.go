package leet_code

import (
	"strings"
	"testing"
)

//分割字符串的最大得分
func maxScore(s string) int {
	var end int
	for i := 1; i < len(s); i++ {
		r := strings.Count(s[:i], "0") + strings.Count(s[i:], "1")
		if end < r {
			end = r
		}
	}
	return end
}

func Test_maxScore(t *testing.T) {
	t.Log(maxScore("11"))
}
