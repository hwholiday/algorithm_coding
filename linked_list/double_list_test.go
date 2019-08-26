package linked_list

import "testing"

func TestNewDoubleList(t *testing.T) {
	l := NewDoubleList()
	l.Append(&DoubleNode{
		Data: 1,
	})
	l.Append(&DoubleNode{
		Data: 5,
	})
	l.Append(&DoubleNode{
		Data: 10,
	})
	t.Log(l.GetAll())
	t.Log(l.GetAllV2())
	t.Log(l.Get(3))
	t.Log(l.InsertNext(l.Get(2), &DoubleNode{
		Data: 11,
	}))

	t.Log(l.InsertNext(l.Get(0), &DoubleNode{
		Data: 3,
	}))
	t.Log(l.GetAll())
	t.Log(l.InsertPrev(l.Get(0), &DoubleNode{
		Data: 0,
	}))
	t.Log(l.GetAll())
	t.Log(l.InsertPrev(l.Get(1), &DoubleNode{
		Data: 2,
	}))
	t.Log(l.GetAll())
}
