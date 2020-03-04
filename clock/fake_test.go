package clock_test

import (
	"testing"
	"time"

	"github.com/braddle/go-http-template/clock"

	"github.com/stretchr/testify/suite"
)

type FakeClockSuite struct {
	suite.Suite
}

func TestFakeClockSuite(t *testing.T) {
	suite.Run(t, new(FakeClockSuite))
}

func (s *FakeClockSuite) TestFakeClockReturnsGivenTime() {
	then := time.Date(1999, 12, 31, 23, 59, 59, 0, time.UTC)
	c := clock.Fake(then)

	s.Equal(then, c.GetCurrentTime())
}
