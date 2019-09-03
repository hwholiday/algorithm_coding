package leet_code

import (
	"fmt"
	"testing"
)

//整数转罗马数字
func TestIntToRoman(t *testing.T) {
	t.Log(intToRoman(1994))
}

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var roman = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var index = 0
	var result string
	for index < len(roman) {
		for num >= nums[index] {
			result += roman[index]
			num = num - nums[index]
			fmt.Println(num)
		}
		index++
	}
	return result
}
