package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

type MinHeap []int

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MinHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func TestMinHeap(t *testing.T) {
	m := &MinHeap{
		2, 10, 5, 9, 12, 7, 3,
	}
	heap.Init(m)
	for m.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(m))
	}

}
