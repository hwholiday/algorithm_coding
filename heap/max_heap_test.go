package heap

import (
	"container/heap"
	"fmt"
	"testing"
)

type MaxHeap []int

func (m MaxHeap) Len() int {
	return len(m)
}

func (m MaxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m MaxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[0 : n-1]
	return x
}

func TestMaxHeap(t *testing.T) {
	m := &MaxHeap{
		2, 10, 5, 9, 12, 7, 3,
	}
	heap.Init(m)
	for m.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(m))
	}

}
