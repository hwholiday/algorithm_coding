package leet_code

import (
	"testing"
)

//最小覆盖子串 暴力破解
func minWindow(s string, t string) string {
	var (
		tMap = make(map[rune]int, len(t))
	)
	for _, v := range t {
		tMap[v]++
	}
	check := func(d string) bool {
		var (
			a = make(map[rune]int, len(d))
		)
		for _, v := range d {
			a[v]++
		}
		for k, v := range tMap {
			val, ok := a[k]
			if !ok || val < v {
				return false
			}
		}
		return true
	}
	var (
		end    = s
		isFind = false
	)
	for i, _ := range s {
		for j, _ := range s {
			if j >= i {
				temp := s[i : j+1]
				if check(temp) {
					isFind = true
					if len(end) > len(temp) {
						end = temp
					}
				}
			}
		}
	}
	if !isFind {
		return ""
	}
	return end
}
func minWindowV2(s string, t string) string {
	var (
		tMap = make(map[string]int, len(t))
		sMap = make(map[string]int)
	)
	for _, v := range t {
		tMap[string(v)]++
	}
	check := func() bool {
		for k, v := range tMap {
			val, ok := sMap[k]
			if !ok || val < v {
				return false
			}
		}
		return true
	}
	var (
		end    = s
		endNum = len(s)
		isFind = false
	)
	for l, r := 0, 0; r < len(s); r++ {
		sMap[string(s[r])]++
		if check() {
			isFind = true
			for check() {
				sMap[string(s[l])]--
				l++
			}
			if endNum > (r+1)-(l-1) {
				endNum = (r + 1) - (l - 1)
				end = s[l-1 : r+1]
			}
		}
	}
	if !isFind {
		return ""
	}
	return end
}

func Test_minWindow(t *testing.T) {
	t.Log(minWindowV2("cabwefgewcwaefgcf", "cae"))
}
