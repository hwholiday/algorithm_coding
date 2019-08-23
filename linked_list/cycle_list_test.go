package linked_list

import "testing"

func TestNewCycleList(t *testing.T) {
	l := NewCycleList()
	l.Append(&CycleNode{
		Data: 1,
	})
	l.Append(&CycleNode{
		Data: 2,
	})
	l.Append(&CycleNode{
		Data: 3,
	})
	t.Log(l.GetAll())
	t.Log(l.Remove(l.Get(1)))
	t.Log(l.GetAll())
}
