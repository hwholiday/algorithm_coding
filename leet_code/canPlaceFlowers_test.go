package leet_code

import (
	"testing"
)

//605. 种花问题
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n > len(flowerbed) {
		return false
	}
	var count int
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return n <= count
}

func Test_canPlaceFlowers(t *testing.T) {
	t.Log(canPlaceFlowers([]int{1, 0}, 1))
}
