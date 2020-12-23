package leet_code

import "testing"

func firstUniqChar(s string) int {
	var (
		raneMap = make(map[rune]int, len(s))
	)
	for _, v := range s {
		raneMap[v]++
	}
	for i, v := range s {
		if val, ok := raneMap[v]; ok {
			if val == 1 {
				return i
			}
		}
	}
	return -1
}

func Test_firstUniqChar(t *testing.T) {
	t.Log(firstUniqChar("loveleetcode"))
}
