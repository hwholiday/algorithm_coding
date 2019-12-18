package wheel_timer

import (
	"time"
)

type SimpleTimingWheel struct {
	interval    time.Duration
	nodeNum     int
	currentNode int
	ticker      *time.Ticker
	list        []*DoubleList
	stop        chan bool
}

type ModOption func(option *SimpleTimingWheel)

func SetInterval(interval time.Duration) ModOption {
	return func(option *SimpleTimingWheel) {
		option.interval = interval
	}
}

func SetNodeNum(nodeNum int) ModOption {
	return func(option *SimpleTimingWheel) {
		option.nodeNum = nodeNum
	}
}

func NewSimpleTimingWheel(option ...ModOption) *SimpleTimingWheel {
	op := &SimpleTimingWheel{
		interval:    time.Second,
		nodeNum:     60 * 60 * 24,
		currentNode: 0,
		list:        make([]*DoubleList, 60*60*24),
		stop:        make(chan bool, 1),
	}
	for _, fn := range option {
		fn(op)
	}
	op.init()
	return op
}
func (s *SimpleTimingWheel) init() {
	for i := 0; i < s.nodeNum; i++ {
		s.list[i] = NewDoubleList()
	}
}

func (s *SimpleTimingWheel) Start() {
	s.ticker = time.NewTicker(s.interval)
	go s.distribution()
}

func (s *SimpleTimingWheel) Stop() {
	s.stop <- true
}

func (s *SimpleTimingWheel) distribution() {
	for {
		select {
		case <-s.ticker.C:
			s.run()
		case <-s.stop:
			s.ticker.Stop()
		}
	}
}

func (s *SimpleTimingWheel) run() {
	task := s.list[s.currentNode]
	s.runTask(task)
	if s.currentNode == s.nodeNum-1 {
		s.currentNode = 0
	} else {
		s.currentNode++
	}
}

func (s *SimpleTimingWheel) runTask(node *DoubleList) {
	if node == nil {
		return
	}
	for _, v := range node.GetAll() {
		go v.f()
	}
}

func (s *SimpleTimingWheel) AddTask(d time.Duration, f func()) {
	delaySeconds := int(d.Seconds())
	s.list[delaySeconds-1].Append(Data{
		f:       f,
		isCycle: false,
	})
}

func (s *SimpleTimingWheel) getPositionAndCircle(d time.Duration) (pos int, circle int) {
	delaySeconds := int(d.Seconds())
	intervalSeconds := int(s.interval.Seconds())
	circle = int(delaySeconds / intervalSeconds / s.nodeNum) //添加轮数
	//计算要放的节点

	pos = int(s.currentNode+delaySeconds/intervalSeconds) % s.nodeNum //计算位置
	return
}
