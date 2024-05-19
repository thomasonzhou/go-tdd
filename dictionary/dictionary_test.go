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
		assertError(err, ErrNotFound, t)
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	dict.Add("日本語", "The japanese language")
	assertDefinition(dict, "日本語", "The japanese language", t)

}

func assertDefinition(dict Dictionary, word, definition string, t *testing.T) {
	got, err := dict.Search(word)
	if err != nil {
		t.Fatal("should find word but didn't", err)
	}
	assertStringsEqual(definition, got, t)
}

func assertError(err, want error, t *testing.T) {
	t.Helper()
	if err != want {
		t.Errorf("expected an error but didn't get one")
	}
}

func assertStringsEqual(s1, s2 string, t *testing.T) {
	t.Helper()
	if s1 != s2 {
		t.Errorf("got %q want %q", s1, s2)
	}
}
