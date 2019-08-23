package linked_list

import (
	"sync"
)

type DoubleNode struct {
	Data interface{}
	Prev *DoubleNode
	Next *DoubleNode
}

type DoubleList struct {
	Size int
	mu   *sync.Mutex
	Head *DoubleNode
	Tail *DoubleNode
}

func NewDoubleList() *DoubleList {
	var l = new(DoubleList)
	l.Size = 0
	l.Head = nil
	l.Tail = nil
	l.mu = new(sync.Mutex)
	return l
}

func (l *DoubleList) Append(node *DoubleNode) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	if node == nil {
		return false
	}
	node.Next = nil
	if l.Size == 0 {
		l.Head = node
		node.Prev = nil
	} else {
		old := l.Tail
		old.Next = node
		node.Prev = old
	}
	l.Tail = node
	l.Size++
	return true
}

func (l *DoubleList) InsertNext(node *DoubleNode, data *DoubleNode) bool {
	if node == nil || data == nil {
		return false
	}
	defer func() {
	}()
	if l.IsTail(node) {
		l.Append(data)
	} else {
		l.mu.Lock()
		defer l.mu.Unlock()
		//取出要插入节点的下一个节点（替换该节点）
		old := node.Next
		old.Prev = data

		data.Prev = node
		data.Next = old

		node.Next = data
		l.Size++
	}
	return true
}
func (l *DoubleList) InsertPrev(node *DoubleNode, data *DoubleNode) bool {
	if node == nil || data == nil {
		return false
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.IsHead(node) {
		old := node
		old.Prev = data
		data.Prev = nil
		data.Next = old
		l.Head = data
	} else {
		//拿当前节点上一个节点数据
		old := node.Prev
		old.Next = data
		data.Prev = old
		data.Next = node
		node.Prev = data
	}
	l.Size++
	return true
}

func (l *DoubleList) Get(i int) *DoubleNode {
	if i >= l.Size || i < 0 {
		return nil
	}
	node := l.Head
	for j := 0; j < i; j++ {
		node = node.Next
	}
	return node
}

func (l *DoubleList) GetAll() []interface{} {
	var data []interface{}
	node := l.Head
	for {
		data = append(data, node.Data)
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return data
}

func (l *DoubleList) GetHead() *DoubleNode {
	return l.Head
}

func (l *DoubleList) GetTail() *DoubleNode {
	return l.Tail
}

func (l *DoubleList) IsHead(node *DoubleNode) bool {
	if l.GetHead() == node {
		return true
	}
	return false
}

func (l *DoubleList) IsTail(node *DoubleNode) bool {
	if l.GetTail() == node {
		return true
	}
	return false
}
