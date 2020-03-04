package clock_test

import (
	"testing"
	"time"

	"github.com/braddle/go-http-template/clock"
	"github.com/stretchr/testify/suite"
)

type ClockSuite struct {
	suite.Suite
}

func TestClockSuite(t *testing.T) {
	suite.Run(t, new(ClockSuite))
}

func (s *ClockSuite) TestReturnsNewTimeAsRequested() {
	c := clock.New()
	t := c.GetCurrentTime()

	s.True(t.Round(time.Millisecond).Equal(time.Now().Round(time.Millisecond)))
}
