package leet_code

import (
	"sort"
	"testing"
)

type myArray [][]int

func (m myArray) Len() int {
	return len(m)
}
func (m myArray) Less(i, j int) bool {
	return m[i][0] > m[j][0]
}

func (m myArray) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

//对角线遍历 II
func findDiagonalOrder(nums myArray) []int {
	var indexMap = make(map[int]myArray)
	for i, v := range nums {
		for j, _ := range v {
			indexMap[i+j] = append(indexMap[i+j], []int{i, j})
		}
	}
	for k, v := range indexMap {
		sort.Sort(v)
		indexMap[k] = v
	}
	var numArray []int
	var indexArray []int
	for k, _ := range indexMap {
		indexArray = append(indexArray, k)
	}
	sort.Ints(indexArray)
	for _, v := range indexArray {
		for _, v1 := range indexMap[v] {
			numArray = append(numArray, nums[v1[0]][v1[1]])
		}
	}
	return numArray
}

func Test_findDiagonalOrder(t *testing.T) {
	var nums myArray
	nums = make([][]int, 3)
	nums[0] = []int{1, 2, 3}
	nums[1] = []int{4, 5, 6}
	nums[2] = []int{7, 8, 9}
	t.Log(findDiagonalOrder(nums))
}
