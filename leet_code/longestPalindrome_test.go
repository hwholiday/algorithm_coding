package leet_code

import (
	"testing"
)

//最长回文子串
func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("ac"))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var info = make(map[string]int)
	for i, v := range s {
		v1 := string(v)
		for j := len(s); j > 0; j-- {
			if v1 == s[j-1:j] {
				if i < j-1 {
					info[s[i:j]] = len(s[i:j])
				}
			}
		}
	}
	var maxLen int
	var key string
	for k, v := range info {
		var cunUse = true
		j := len(k)
		for i := 0; i < len(k); i++ {
			if k[i:i+1] != k[j-1:j] {
				//fmt.Println("not ", k, k[i:i+1], k[j-1:j], j-1, j)
				cunUse = false
				break
			}
			j--
		}
		if cunUse {
			if v > maxLen {
				maxLen = v
				key = k
			}
		}
	}
	if key == "" {
		return s[:1]
	}
	return key
}
