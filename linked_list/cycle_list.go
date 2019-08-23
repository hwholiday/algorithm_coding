package linked_list

import (
	"sync"
)

type CycleNode struct {
	Data interface{}
	Next *CycleNode
}

type CycleList struct {
	Size int
	Mu   *sync.Mutex
	Head *CycleNode
}

func NewCycleList() *CycleList {
	return &CycleList{
		Size: 0,
		Mu:   new(sync.Mutex),
		Head: nil,
	}
}

func (l *CycleList) Append(node *CycleNode) bool {
	if node == nil {
		return false
	}
	l.Mu.Lock()
	defer l.Mu.Unlock()
	if l.Size == 0 {
		l.Head = node
	} else {
		old := l.Head
		//找到尾部
		for {
			if old.Next == l.Head {
				break
			}
			old = old.Next
		}
		old.Next = node
	}
	node.Next = l.Head //新添加上的链接上头部
	l.Size++
	return true
}

func (l *CycleList) GetAll() []interface{} {
	var data []interface{}
	old := l.Head
	for {
		data = append(data, old.Data)
		//当前的下一个节点是头部认为到了尾部
		if old.Next == l.Head {
			break
		}
		old = old.Next
	}
	return data
}

func (l *CycleList) Get(i int) *CycleNode {
	if i < 0 || i >= l.Size {
		return nil
	}
	old := l.Head
	for j := 0; j < i; j++ {
		old = old.Next
	}
	return old
}

func (l *CycleList) Remove(node *CycleNode) bool {
	if node == nil {
		return false
	}
	//找到该节点
	old := l.Head
	for {
		if old.Next == node { //找到该节点前一个节点
			break
		}
		old = old.Next
	}
	old.Next = node.Next //去掉node的数据
	return true
}
