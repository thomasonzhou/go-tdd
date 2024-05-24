package context

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	// "time"
)

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, char := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(char)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

func TestServer(t *testing.T) {
	t.Run("strings equal", func(t *testing.T) {
		data := "Bonjour, 世界"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q wanted %q", response.Body.String(), data)
		}
		// store.assertStoreWasCancelled()
	})

	// t.Run("cancel if Cancel called", func(t *testing.T) {
	// 	data := "houston has a problem"
	// 	store := &SpyStore{response: data, t: t}
	// 	server := Server(store)

	// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
	// 	cancellingCtx, cancel := context.WithCancel(request.Context())
	// 	time.AfterFunc(7*time.Millisecond, cancel)
	// 	request = request.WithContext(cancellingCtx)

	// 	response := httptest.NewRecorder()

	// 	server.ServeHTTP(response, request)

	// 	store.assertStoreCancelled()
	// })
}

// func (s *SpyStore) assertStoreCancelled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Error("expected store to be cancelled")
// 	}
// }

// func (s *SpyStore) assertStoreWasCancelled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Error("store is not supposed to be cancelled")
// 	}
// }
