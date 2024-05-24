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
		store := &SpyStore{data, false}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got %q wanted %q", response.Body.String(), data)
		}
		assertStoreWasCancelled(store, t)

	})

	t.Run("cancel if Cancel called", func(t *testing.T) {
		data := "houston has a problem"
		store := &SpyStore{response: data}
		server := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(7*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStoreCancelled(store, t)
	})
}

func assertStoreCancelled(store *SpyStore, t *testing.T) {
	if !store.cancelled {
		t.Error("expected store to be cancelled")
	}
}

func assertStoreWasCancelled(store *SpyStore, t *testing.T) {
	t.Helper()
	if store.cancelled {
		t.Error("store is not supposed to be cancelled")
	}
}
