package wheel_timer

import "time"

type SimpleTimingWheel struct {
}

type ModOption func(option *SimpleTimingWheel)

func NewSimpleTimingWheel(option ...ModOption) *SimpleTimingWheel {
	op := &SimpleTimingWheel{}
	for _, fn := range option {
		fn(op)
	}
	return op
}

func (s *SimpleTimingWheel) AddTask(duration time.Duration, f func()) {

}
