package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(111 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("strings equal", func(t *testing.T) {
		data := "Bonjour, 世界"
		store := &SpyStore{data, false, t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q wanted %q", response.Body.String(), data)
		}
		store.assertStoreWasCancelled()
	})

	t.Run("cancel if Cancel called", func(t *testing.T) {
		data := "houston has a problem"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(7*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		store.assertStoreCancelled()
	})
}

func (s *SpyStore) assertStoreCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("expected store to be cancelled")
	}
}

func (s *SpyStore) assertStoreWasCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store is not supposed to be cancelled")
	}
}
