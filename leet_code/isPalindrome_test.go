package leet_code

import (
	"strconv"
	"testing"
)

//回文数
func TestIsPalindrome(t *testing.T) {
	t.Log(isPalindrome(10))

}
func isPalindrome(x int) bool {
	sX := strconv.Itoa(x)
	var j = len(sX)
	for _, v := range sX {
		if string(v) != sX[j-1:j] {
			return false
		}
		j--
	}
	return true
}
