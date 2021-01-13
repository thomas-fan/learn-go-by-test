package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const sleep = "sleep"
const write = "write"

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}


type SpyTime struct {
	durationSlept time.Duration
}

func (c *ConfigurableSleeper) Sleep()  {
	c.sleep(c.duration)
}

func (s *SpyTime) Sleep(duration time.Duration)  {
	s.durationSlept = duration
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct {
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}
