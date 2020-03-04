package clock

import "time"

type FakeClock struct {
	t time.Time
}

func (f FakeClock) GetCurrentTime() time.Time {
	return f.t
}

func Fake(t time.Time) CurrentTimeFactory {
	return FakeClock{t}
}
