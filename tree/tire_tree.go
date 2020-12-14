package tree

type Trie struct {
	node   map[rune]*Trie
	val    string
	isLast bool
}

func Constructor() Trie {
	return Trie{node: make(map[rune]*Trie, 26), isLast: false}
}

func (t *Trie) Insert(word string) {
	for _, v := range word {
		if t.node[v] == nil { //没找到该节点
			t.node[v] = &Trie{node: make(map[rune]*Trie, 26), isLast: false}
		}
		t = t.node[v] //找到节点，跳到下一个节点
	}
	t.val = word
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

func (t *Trie) StartsPrx(prefix string) (res *[]interface{}) {
	for _, v := range prefix {
		if t.node[v] == nil { //没找到该节点
			return
		}
		t = t.node[v] //找到节点，跳到下一个节点
	}
	res = &[]interface{}{}
	GetStartsPrx(t, res)
	return
}

func GetStartsPrx(t *Trie, res *[]interface{}) {
	if t.isLast {
		*res = append(*res, t.val)
	}
	for _, v := range t.node {
		GetStartsPrx(v, res)
	}
	return
}

//查找数据
func (t *Trie) SearchNode(key string) (res []interface{}) {
	root := t
	for _, v := range key {
		if v, ok := root.node[v]; ok {
			root.node = v.node
		} else {
			return
		}
	}
	var queue []*Trie
	queue = append(queue, root)
	for len(queue) > 0 {
		var childQueue []*Trie
		for _, v := range queue { //遍历这层的数据
			if v.isLast {
				res = append(res, v.val)
			}
			for _, vi := range v.node { //把下一次的数据进行遍历
				childQueue = append(childQueue, vi)
			}
		}
		queue = childQueue
	}
	return
}
