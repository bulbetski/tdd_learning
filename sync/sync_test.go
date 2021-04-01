package sync

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	t.Run("incrementing counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Incr()
		counter.Incr()
		counter.Incr()

		assertCounter(t, counter, 3)
	})

	t.Run("counter runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Incr()
				w.Done()
			}(&wg)
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, c *Counter, want int) {
	t.Helper()
	got := c.Value()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
