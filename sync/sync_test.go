package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("增加 3 次后,值为 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})
	t.Run("并发安全计数", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()
		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t *testing.T, got *Counter, want int)  {
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}