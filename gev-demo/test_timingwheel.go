package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/RussellLuo/timingwheel"
)

type EveryScheduler struct {
	Interval time.Duration
}

func (s *EveryScheduler) Next(prev time.Time) time.Time {
	return prev.Add(s.Interval)
}

func main() {
	fmt.Println(runtime.GOOS)
	tw := timingwheel.NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	exitC := make(chan time.Time)
	t := tw.ScheduleFunc(&EveryScheduler{time.Second}, func() {
		fmt.Println("The timer fires")
		exitC <- time.Now().UTC()
	})

	<-exitC
	<-exitC

	// We need to stop the timer since it will be restarted again and again.
	for !t.Stop() {
	}

}
