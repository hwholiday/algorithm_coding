package leet_code

import (
	"strconv"
	"testing"
)

//棒球比赛
func calPoints(ops []string) int {
	array := make([]int, 0, len(ops))
	for _, v := range ops {
		switch v {
		case "+":
			array = append(array, array[len(array)-1]+array[len(array)-2])
		case "D":
			array = append(array, array[len(array)-1]*2)
		case "C":
			array = array[:len(array)-1]
		default:
			i, _ := strconv.Atoi(v)
			array = append(array, i)
		}
	}
	var num int
	for _, v := range array {
		num += v
	}
	return num
}

func Test_calPoints(t *testing.T) {
	ps := []string{"5", "2", "C", "D", "+"}
	t.Log(calPoints(ps))
}
