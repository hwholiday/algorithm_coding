package leet_code

import (
	"testing"
)

//罗马数字转整数
func TestRomanToInt(t *testing.T) {
	t.Log(romanToInt("III"))

}

func romanToInt(s string) int {
	var (
		result int
		nums   []int
	)
	var data = make(map[string]int)
	data["M"] = 1000
	data["CM"] = 900
	data["D"] = 500
	data["CD"] = 400
	data["C"] = 100
	data["XC"] = 90
	data["L"] = 50
	data["XL"] = 40
	data["X"] = 10
	data["IX"] = 9
	data["V"] = 5
	data["IV"] = 4
	data["I"] = 1
	for j := 0; j < len(s); {
		if j+2 <= len(s) {
			if v1, ok := data[s[j:j+2]]; ok {
				nums = append(nums, v1)
				j = j + 2
			} else if v, ok := data[s[j:j+1]]; ok {
				nums = append(nums, v)
				j++
			}
		} else {
			if v, ok := data[s[j:j+1]]; ok {
				nums = append(nums, v)
				j++
			}
		}
	}
	for _, v := range nums {
		result += v
	}
	return result
}
