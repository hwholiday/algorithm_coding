package leet_code

import (
	"testing"
)

//134. 加油站
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(cost); i++ {
		var cur int = 0
		var mGas = getList(gas, i)
		var mCost = getList(cost, i)
		for j := 0; j < len(mGas); j++ {
			if j-1 < 0 {
				cur = cur + mGas[j]
			} else {
				if cur-mCost[j-1] < 0 {
					cur = 0
					break
				}
				cur = cur + mGas[j] - mCost[j-1]
			}
			if cur <= 0 {
				cur = 0
				break
			}
			if j+1 == len(mGas) {
				if cur >= mCost[j] {
					return i
				}
			}
		}
	}
	return -1
}

func getList(gas []int, i int) (data []int) {
	data = append(gas[i:], gas[:i]...)
	return
}

func Test_canCompleteCircuit(t *testing.T) {
	var (
		gas  = []int{1, 2, 3, 4, 5}
		cost = []int{3, 4, 5, 1, 2}
	)
	t.Log(canCompleteCircuit(gas, cost))
	t.Log(getList(gas, 3))
}
