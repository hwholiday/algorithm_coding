package leet_code

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Log(lengthOfLongestSubstring("abcdefgafkzwb"))
}

func lengthOfLongestSubstring(s string) int {
	lastOccurred := make(map[byte]int)
	start := 0
	maxLength := 0
	for i, ch := range []byte(s) {
		//找到上次该字母出现的起点
		//判断该位置是否大于起点
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1 //从该字母后一位作为起点,以前的存在相同的而且已经判断就不再判断
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
