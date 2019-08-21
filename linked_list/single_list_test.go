package linked_list

import (
	"testing"
)

func TestNew(t *testing.T) {
	l := New()
	l.Append(&Node{
		data: "1",
	})
	l.Append(&Node{
		data: "2",
	})
	l.Append(&Node{
		data: "3",
	})
	t.Log(l.GetAll())
	t.Log(l.Get(2))
	l.Remove(1)
	t.Log(l.GetAll())
	l.Clear()
	t.Log(l.GetAll())
}
