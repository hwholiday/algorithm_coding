package wheel_timer

import "testing"

func TestNewDoubleList(t *testing.T) {
	l := NewDoubleList()
	l.Append(Data{isCycle: false, Id: 1, f: func() {
		t.Log(1)
	}})
	l.Append(Data{isCycle: true, Id: 2, f: func() {
		t.Log(2)
	}})
	l.Append(Data{isCycle: false, Id: 3, f: func() {
		t.Log(3)
	}})
	l.Append(Data{isCycle: true, Id: 4, f: func() {
		t.Log(4)
	}})
	l.Append(Data{isCycle: false, Id: 5, f: func() {
		t.Log(5)
	}})
	t.Log(l.GetAll())
	t.Log(l.GetAll())
}
