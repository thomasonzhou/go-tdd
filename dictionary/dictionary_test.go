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
	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Add("日本語", "The japanese language")
		assertError(err, nil, t)
		assertDefinition(dict, "日本語", "The japanese language", t)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "Français"
		definition := "The french language"
		dict := Dictionary{word: definition}
		err := dict.Add(word, "Le langue des francophones")
		assertError(err, ErrWordExists, t)
		assertDefinition(dict, word, definition, t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "りんご"
		newDefinition := "Apple in japanese"
		dict := Dictionary{word: "赤い果物"}
		err1 := dict.Update(word, newDefinition)
		assertError(err1, nil, t)

		definition, err2 := dict.Search(word)
		assertError(err2, nil, t)
		assertStringsEqual(newDefinition, definition, t)
	})

	t.Run("new word", func(t *testing.T) {
		dict := Dictionary{}

		err := dict.Update("doesn't exist", "doesn't matter")
		assertError(err, ErrWordDoesNotExist, t)
	})
}

func TestDelete(t *testing.T) {
	word := "golang"
	dict := Dictionary{word: "a fun language made by people at Google"}
	dict.Delete(word)

	_, err := dict.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}

}

func assertDefinition(dict Dictionary, word, definition string, t *testing.T) {
	t.Helper()
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
