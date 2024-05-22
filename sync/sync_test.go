package sync

import "testing"

func TestSync(t *testing.T) {
	t.Run("count to 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value != 3 {
			t.Errorf("got %d wanted %d", counter.Value, 3)
		}
	})

}
