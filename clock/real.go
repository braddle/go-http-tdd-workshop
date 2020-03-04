package clock

import "time"

type CurrentTimeFactory interface {
	GetCurrentTime() time.Time
}

type Clock struct{}

func (c Clock) GetCurrentTime() time.Time {
	return time.Now()
}

func New() CurrentTimeFactory {
	return Clock{}
}
