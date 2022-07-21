package stopwatch

import (
	"fmt"
	"time"
)

type StopWatch struct {
	start time.Time
}

func Make() *StopWatch {
	return &StopWatch{}
}

func (s *StopWatch) Start()  {
	s.start = time.Now()
}

func (s *StopWatch) Lap() {
	fmt.Printf("%fsec\n", (time.Now().Sub(s.start)).Seconds())
}