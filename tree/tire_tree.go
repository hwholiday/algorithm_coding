package tree

import "fmt"

type Trie struct {
	node   map[rune]*Trie
	isLast bool
}

func Constructor() Trie {
	return Trie{node: make(map[rune]*Trie, 26), isLast: false}
}

func (t *Trie) Insert(word string) {
	for _, v := range word {
		if t.node[v] == nil { //没找到该节点
			vt := &Trie{node: make(map[rune]*Trie, 26), isLast: false}
			t.node[v] = vt
		}
		t = t.node[v] //找到节点，跳到下一个节点
	}
	t.isLast = true
}

func (t *Trie) Search(word string) bool {
	for _, v := range word {
		if t.node[v] == nil { //没找到该节点
			return false
		}
		t = t.node[v] //找到节点，跳到下一个节点
	}
	return t.isLast
}

func (t *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if t.node[v] == nil { //没找到该节点
			return false
		}
		t = t.node[v] //找到节点，跳到下一个节点
	}
	return true
}

func (t *Trie) GetAll(prefix string) *[]interface{} {
	data := new([]interface{})
	node := map[rune]*Trie{}
	for _, v := range prefix { //找到全匹配下的节点
		if t.node[v] != nil {
			if t.node[v].isLast { //判断当前节点是不是一个完整单词
				*data = append(*data, prefix)
			}
			t = t.node[v] //跳到下一个节点
			node = t.node

		}
	}
	getNodeAll(prefix, node, data)
	return data
}

var a string

func getNodeAll(prefix string, node map[rune]*Trie, data *[]interface{}) {
	fmt.Println("---------------------- prefix ", prefix)
	for k, v := range node {
		fmt.Println("---------------------- prefix  v1 ", prefix)

		prefix = prefix + string(k)
		fmt.Println("---------------------- prefix v2 ", prefix)
		//fmt.Println(string(k),prefix)
		if v.isLast {
			fmt.Println("---------------------- 获得结果", prefix)
			*data = append(*data, prefix)
		}
		if v.node != nil {
			fmt.Println("---------------------- getNodeAll", prefix)
			getNodeAll(prefix, v.node, data)
		}
	}
}
