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

	fs := fstest.MapFS{
		"how_to_live.md":    {Data: []byte("Title: Juggle life")},
		"global_warming.md": {Data: []byte("Title: The planet is warming up.")},
	}

	posts, err := blogposts.NewBlogPost(fs)

	if err != nil {
		t.Fatal("got error but expected none")
	}

	if len(fs) != len(posts) {
		t.Errorf("want %v, got %v", len(fs), len(posts))
	}

	got := posts[0]
	// the reason why this is first is go maps are ordered by key
	want := blogposts.Post{Title: "The planet is warming up."}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v, got %v", want, got)
	}

}
