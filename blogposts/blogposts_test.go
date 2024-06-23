package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/thomasonzhou/go-tdd/blogposts"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("You learn faster if you fail fast.")
}

func TestNewBlogPost(t *testing.T) {
	const (
		firstBody = `Title: Juggle life
Description: Do things that you like that are difficult
Tags: shower-thoughts
---
I like this advice because everyone has their own version of it.
"Do what excites"
"I took the one less traveled by, And that has made all the difference."
"Just do it"
`
		secondBody = `Title: The planet is warming up
Description: We should bring less things on the planes
Tags: climate-change, shower-thoughts
---
What if could fit everything in a carry-on?
Is it not enough space?
`
	)

	fs := fstest.MapFS{
		"how_to_live.md":    {Data: []byte(firstBody)},
		"global_warming.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewBlogPost(fs)

	if err != nil {
		t.Fatal("got error but expected none")
	}

	if len(fs) != len(posts) {
		t.Errorf("want %v, got %v", len(fs), len(posts))
	}

	// the reason why this is first is go maps are ordered by key
	want := blogposts.Post{
		Title:       "The planet is warming up",
		Description: "We should bring less things on the planes",
		Tags:        []string{"climate-change", "shower-thoughts"},
		Body: `What if could fit everything in a carry-on?
Is it not enough space?`,
	}
	got := posts[0]
	assertPost(t, want, got)

}

func assertPost(t *testing.T, want blogposts.Post, got blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v got %v", want, got)
	}
}
