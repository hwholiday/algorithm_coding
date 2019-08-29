package leet_code

import (
	"math"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

//字符串转换整数 (atoi)
func TestMyAtoi(t *testing.T) {
	t.Log(myAtoi("   +12"))
	t.Log(myAtoi(" ##2  -12"))
	t.Log(myAtoi("   -12"))
	t.Log(myAtoi(" 2  -12"))
	t.Log(myAtoi("-"))
	t.Log(myAtoi("-91283472332"))
	t.Log(myAtoi("3.114"))
	t.Log(myAtoi("+-2"))
	t.Log(myAtoi("   +1 123"))
	t.Log(myAtoi("  0000000000012345678"))
	t.Log(myAtoi("0-1"))
	t.Log(myAtoi("-42"))
	t.Log(myAtoi("42"))
	t.Log(myAtoi("-   234"))
	t.Log(myAtoi("words and 987"))
	t.Log(myAtoi("4193 with words"))
	t.Log(myAtoi(".1"))
}

func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	str = strings.Trim(str, " ")
	var (
		min    = math.Pow(-2, 31)
		max    = math.Pow(2, 31) - 1
		start  int
		sign   float64 = 1
		result float64
		accept string
		err    error
	)
	for i, v := range str {
		if !unicode.IsSpace(v) {
			start = i
			if unicode.IsDigit(v) { //第一位是数字要包含
				start = i - 1
			} else {
				if strings.Contains(string(v), "-") {
					sign = -1
				} else if strings.Contains(string(v), "+") {
					sign = 1
				} else {
					return 0
				}
			}
			break
		}
	}
	for i, v := range str {
		if i <= start {
			continue
		}
		if !unicode.IsDigit(v) {
			break
		}
		accept += string(v)
	}
	result, err = strconv.ParseFloat(accept, 64)
	if err != nil {
		return 0
	}
	result = result * sign
	if result > max {
		return int(max)
	}
	if result < min {
		return int(min)
	}
	return int(result)
}
