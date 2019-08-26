package leet_code

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcabcbb"))
}

type Data struct {
	Index int
	Val   int
	have  bool
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	return 0
}
