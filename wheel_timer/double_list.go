package wheel_timer

import (
	"sync"
)

type Data struct {
	circle int
	f      func()
}

type Node struct {
	Data Data
	Prev *Node
	Next *Node
}

type DoubleList struct {
	size uint32
	head *Node
	tail *Node
	mu   *sync.Mutex
}

func NewDoubleList() *DoubleList {
	return &DoubleList{
		size: 0,
		head: nil,
		tail: nil,
		mu:   new(sync.Mutex),
	}
}
func (l *DoubleList) Append(data Data) {
	l.mu.Lock()
	defer l.mu.Unlock()
	node := &Node{
		Data: data,
		Prev: nil,
		Next: nil,
	}
	if l.size == 0 {
		l.head = node
	} else {
		old := l.tail
		old.Next = node
		node.Prev = old
	}
	l.tail = node
	l.size++
}

func (l *DoubleList) GetAll() (data []Data) {
	node := l.head
	for node != nil {
		data = append(data, node.Data)
		l.mu.Lock()
		if node.Data.circle > 0 {
			node.Data.circle--
		} else {
			l.remove(node)
		}
		l.mu.Unlock()
		node = node.Next
	}
	return data
}

func (l *DoubleList) remove(node *Node) {
	//一次性任务,删除
	if l.IsHead(node) {
		l.head = node.Next
	} else if l.IsTail(node) {
		old := node.Prev
		old.Next = nil
		l.tail = node

	} else {
		old := node.Prev
		old.Next = node.Next
	}
	l.size--
}
func (l *DoubleList) GetHead() *Node {
	return l.head
}
func (l *DoubleList) GetTail() *Node {
	return l.tail
}

func (l *DoubleList) IsHead(node *Node) bool {
	if l.GetHead() == node {
		return true
	}
	return false
}

func (l *DoubleList) IsTail(node *Node) bool {
	if l.GetTail() == node {
		return true
	}
	return false
}
