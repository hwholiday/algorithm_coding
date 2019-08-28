package leet_code

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	t.Log(longestPalindrome("aaabaaaa"))
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
	fmt.Println(info)
	for k, v := range info {
		var cunUse = true
		var endFor = false
		for i := 0; i < len(k); i++ {
			for j := len(k); j > 0; j-- {
				if k[i:i+1] != k[j-1:j] {
					cunUse = false
					endFor = true
					fmt.Println("not ", k, k[i:i+1], k[j-1:j], j-1, j)
					break
				}
			}
			if endFor {
				break
			}
		}
		if cunUse {
			fmt.Println(k)
			if v >= maxLen {
				key = k
			}
		}
	}
	return key
}
