package standard_library

import (
	"sort"
	"testing"
)

type Data struct {
	Id      int
	Weights int
}

type DataArray []Data

func (d DataArray) Len() int {
	return len(d)
}
func (d DataArray) Less(i, j int) bool {
	return d[i].Weights > d[j].Weights
}
func (d DataArray) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func Test_sort(t *testing.T) {
	data := DataArray{
		{Id: 2, Weights: 2},
		{Id: 1, Weights: 1},
		{Id: 9, Weights: 9},
		{Id: 8, Weights: 8},
		{Id: 3, Weights: 3},
		{Id: 6, Weights: 6},
		{Id: 10, Weights: 10},
	}
	t.Log(data)
	sort.Sort(data)
	t.Log(data)
}

func Test_sortSlice(t *testing.T) {
	data := []Data{
		{Id: 2, Weights: 2},
		{Id: 1, Weights: 1},
		{Id: 9, Weights: 9},
		{Id: 8, Weights: 8},
		{Id: 3, Weights: 3},
		{Id: 6, Weights: 6},
		{Id: 10, Weights: 10},
	}
	t.Log(data)
	sort.Slice(data, func(i, j int) bool {
		return data[i].Weights < data[j].Weights
	})
	t.Log(data)
}
