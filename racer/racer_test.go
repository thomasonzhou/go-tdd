package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test faster server", func(t *testing.T) {
		slowServer := getDelayedServer(20 * time.Millisecond)
		fastServer := getDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("return error if server takes more than 10s", func(t *testing.T) {
		server1 := getDelayedServer(23 * time.Millisecond)
		server2 := getDelayedServer(11 * time.Millisecond)

		defer server1.Close()
		defer server2.Close()

		_, err := ConfigurableRacer(server1.URL, server2.URL, 2*time.Millisecond)

		if err == nil {
			t.Error("expected error but got nil")
		}
	})

}

func getDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
