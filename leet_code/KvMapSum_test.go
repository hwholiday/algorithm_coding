package leet_code

import (
	"fmt"
	"testing"
)

//键值映射
type MapSum struct {
	tire *TireData
}

func Constructor() MapSum {
	return MapSum{tire: &TireData{node: make(map[rune]*TireData, 26)}}
}

func (this *MapSum) Insert(key string, val int) {
	this.tire.Insert(key, val)
}

func (this *MapSum) Sum(prefix string) int {
	res := this.tire.Search(prefix)
	if res == nil {
		return 0
	}
	var data int
	for _, v := range res {
		data += v.val
	}
	return data
}

type TireData struct {
	node  map[rune]*TireData
	key   string
	val   int
	isEnd bool
}

func (t *TireData) Insert(key string, val int) {
	for _, v := range key {
		if t.node[v] == nil {
			t.node[v] = &TireData{node: make(map[rune]*TireData, 26)}
		}
		t = t.node[v]
	}
	t.key = key
	t.val = val
	t.isEnd = true
}
func (t *TireData) Search(key string) (data []*TireData) {
	//app
	for _, v := range key {
		if t.node[v] == nil {
			return
		}
		t = t.node[v]
	}
	GetNode(t, &data)
	return data
}

func GetNode(d *TireData, data *[]*TireData) {
	if d.isEnd {
		*data = append(*data, d)
		fmt.Println("d", d.key)
	}
	for _, v := range d.node {
		GetNode(v, data)
	}
}

func Test_TireData(t *testing.T) {
	var tire = &TireData{node: make(map[rune]*TireData, 26)}
	tire.Insert("app", 2)
	tire.Insert("apple", 3)
	tire.Search("app")
}
