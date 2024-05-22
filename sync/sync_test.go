package sync

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	t.Run("count to 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCount(counter, 3, t)
	})

	t.Run("concurrent counting", func(t *testing.T) {
		counter := Counter{}
		wantCount := 999

		var wg sync.WaitGroup
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCount(counter, wantCount, t)
	})

}

func assertCount(counter Counter, want int, t *testing.T) {
	if counter.Value() != want {
		t.Errorf("got %d wanted %d", counter.Value(), want)
	}
}
