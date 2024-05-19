package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"grit": "doing hard stuff for a bit"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("grit")
		want := "doing hard stuff for a bit"
		assertStringsEqual(got, want, t)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := "could not find the word"

		if err == nil {
			t.Fatal("expected an error but didn't get one")
		}
		assertStringsEqual(want, err.Error(), t)
	})
}

func assertStringsEqual(s1, s2 string, t *testing.T) {
	t.Helper()
	if s1 != s2 {
		t.Errorf("got %q want %q", s1, s2)
	}
}
