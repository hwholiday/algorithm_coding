package leet_code

import (
	"testing"
)

//最长公共前缀
func TestLongestCommonPrefix(t *testing.T) {
	t.Log(longestCommonPrefix([]string{"aca", "cba"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var (
		startData = strs[0]
		result    string
	)
	strs = strs[1:]
	var enFor bool
	for i, v := range startData {
		var add bool
		for _, v1 := range strs {
			if i+1 > len(v1) {
				enFor = true
				break
			}
			if string(v) == v1[i:i+1] {
				add = true
			} else {
				add = false
				enFor = true
				break
			}
		}
		if enFor {
			break
		}
		if add {
			result += string(v)
		}
	}
	return result
}
