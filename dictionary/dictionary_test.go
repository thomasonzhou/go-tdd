package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dict := Dictionary{"grit": "doing hard stuff for a bit"}

	got := dict.Search("grit")
	want := "doing hard stuff for a bit"
	assertStringsEqual(got, want, t)
}

func assertStringsEqual(s1, s2 string, t *testing.T) {
	t.Helper()
	if s1 != s2 {
		t.Errorf("got %q want %q", s1, s2)
	}
}
