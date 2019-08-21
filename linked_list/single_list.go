package linked_list

import (
	"fmt"
	"sync"
)

//节点
type Node struct {
	data interface{}
	next *Node
}

type List struct {
	size uint32 //节点总数
	head *Node  //头节点
	tail *Node  //尾节点
	mu   sync.Mutex
}

func New() (l *List) {
	l = new(List)
	l.size = 0
	l.head = nil
	l.tail = nil
	return
}

func (l *List) Append(node *Node) bool {
	if node == nil {
		return false
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	node.next = nil //新增加的节点，下一个节点没有
	if l.size == 0 {
		l.head = node //没数据添加到首节点
	} else {
		old := l.tail //以前老尾部节点指向新添加的节点
		old.next = node
	}
	l.tail = node //成为新的尾部
	l.size++
	return true
}

func (l *List) GetAll() (data []interface{}) {
	if l.head == nil {
		return
	}
	data = append(data, l.head.data) //获取首节点数据
	node := l.head.next              //获取下一个节点数据
	for {
		if node == nil {
			break
		}
		data = append(data, node.data)
		node = node.next //获取下一个节点数据
	}
	return data
}

func (l *List) Get(i uint32) (data interface{}) {
	if i >= l.size {
		return
	}
	info := l.head
	for j := uint32(0); j < i; j++ {
		info = info.next
	}
	return info.data
}

func (l *List) Remove(i uint32) {
	if i >= l.size {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	if i == 0 { //删除头
		l.head = l.head.next
	} else {
		//查询要删除节点的上一个节点
		info := l.head
		for j := uint32(1); j < i; j++ {
			info = info.next
		}
		fmt.Println(info.data)
		node := info.next //要删除的节点
		info.next = node.next
		if i == l.size-1 { //最后一个节点
			l.tail = info
		}
	}
	l.size--
}

func (l *List) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}
