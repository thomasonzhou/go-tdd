package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("custom name in greeting", func(t *testing.T) {
		got := Hello("世界", "")
		want := "Hello, 世界"
		assertStringEqual(t, got, want)
	})
	t.Run("no name prints default greeting", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, 世界"
		assertStringEqual(t, got, want)

	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Picasso", "Spanish")
		want := "Hola, Picasso"
		assertStringEqual(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Robespierre", "French")
		want := "Bonjour, Robespierre"
		assertStringEqual(t, got, want)
	})
}

func assertStringEqual(t testing.TB, s1, s2 string) {
	t.Helper()
	if s1 != s2 {
		t.Errorf("got %q want %q", s1, s2)
	}
}
