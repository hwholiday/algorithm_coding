package main

import "fmt"

type TireNode struct {
	child map[rune]*TireNode
}

type RuneTire struct {
	root *TireNode
}

func NewRuneTire() *RuneTire {
	return &RuneTire{root: &TireNode{
		child: make(map[rune]*TireNode)}}
}

func (r *RuneTire) Insert(key string) {
	node := r.root //从根节点查询
	for _, v := range []rune(key) {
		if node.child[v] == nil { //查询这个字节是否存在根节点下
			//不存在
			node.child[v] = &TireNode{child: make(map[rune]*TireNode)}
		}
		node = node.child[v] //跳到这个节点下
	}
}

func (r *RuneTire) Search(key string) bool {
	node := r.root
	for _, v := range []rune(key) {
		n, ok := node.child[v] //查询这个字节是否存在根节点下
		if !ok {               //不存在
			return false
		}
		node = n //跳到这个节点下
	}
	return true
}

func main() {
	tree := NewRuneTire()
	tree.Insert("abcdef")
	tree.Insert("abe")
	tree.Insert("abf")
	tree.Insert("abe")
	fmt.Println(tree.Search("abe"))
	fmt.Println(tree.Search("abec"))

}
