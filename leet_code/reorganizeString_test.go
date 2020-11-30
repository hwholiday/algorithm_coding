package leet_code

import (
	"testing"
)

//重构字符串
func reorganizeString(S string) string {
	n := len(S)
	if n <= 1 {
		return S
	}
	var vA = [26]int{}
	var maxNum int
	for _, v := range S {
		v -= 'a'
		vA[v]++
		if maxNum < vA[v] {
			maxNum = vA[v]
		}
	}
	if maxNum > (n+1)/2 {
		return ""
	}
	str := make([]byte, n)
	//偶数下标 奇数下标
	evenIdx, oddIdx, halfLen := 0, 1, n/2
	for i, c := range vA {
		ch := byte('a' + i)
		for c > 0 && c <= halfLen && oddIdx < n {
			str[oddIdx] = ch
			c--
			oddIdx += 2
		}
		for c > 0 {
			str[evenIdx] = ch
			c--
			evenIdx += 2
		}
	}
	return string(str)
}

func TestReorganizeString(t *testing.T) {
	reorganizeString("aab")
}
