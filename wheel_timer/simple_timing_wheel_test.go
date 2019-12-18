package wheel_timer

import (
	"fmt"
	"testing"
	"time"
)

func TestNewDoubleList(t *testing.T) {
	l := NewDoubleList()
	l.Append(Data{circle: 1, f: func() {
		t.Log(1)
	}})
	l.Append(Data{circle: 2, f: func() {
		t.Log(2)
	}})
	l.Append(Data{circle: 3, f: func() {
		t.Log(3)
	}})
	l.Append(Data{circle: 4, f: func() {
		t.Log(4)
	}})
	l.Append(Data{circle: 5, f: func() {
		t.Log(5)
	}})
	l.Append(Data{circle: 6, f: func() {
		t.Log(6)
	}})

	t.Log(l.GetAll())
	t.Log(l.GetAll())
}

func TestV1(t *testing.T) {
	delaySeconds := int(7)
	intervalSeconds := int(1)
	circle := int(delaySeconds / intervalSeconds / 6)
	fmt.Println("circle =", int(delaySeconds/intervalSeconds/6)) //添加轮数
	//计算要放的节点
	//剩下的时候在一轮就可以搞定
	fmt.Println("pos =", int(delaySeconds-(intervalSeconds*6*circle)))
}

func TestNewSimpleTimingWheel(t *testing.T) {
	tw := NewSimpleTimingWheel(SetNodeNum(6))
	tw.Start()
	fmt.Println("开始", time.Now().Format("2006-01-02 15:04:05"))
	tw.AddTask(8*time.Second, func() {
		fmt.Println("8秒后执行", time.Now().Format("2006-01-02 15:04:05"))
		//tw.Stop()
	})
	time.Sleep(time.Hour)
}
