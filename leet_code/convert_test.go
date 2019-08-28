package leet_code

import (
	"strings"
	"testing"
)

// Z 字形变换
func TestConvert(t *testing.T) {
	t.Log(convert("LEETCODEISHIRING", 4))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var (
		data  = make([]string, numRows)
		index = 0
	)
	//0   0
	//1  1  1
	//2 2    2
	//3       3
	var isDown = true
	for _, v := range s {
		if isDown {
			data[index] += string(v)
			index++
		} else {
			data[index] += string(v)
			index--

		}
		if index == numRows-1 { //转向
			isDown = false
		} else if index == 0 { //转向
			isDown = true
		}
	}
	return strings.Join(data, "")
}
