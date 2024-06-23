package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	blogrenderer "github.com/thomasonzhou/go-tdd/blogrenderer"
)

func TestBlogRenderer(t *testing.T) {
	var (
		Post = blogrenderer.Post{
			Title:       "How much is too much?",
			Description: "Water consumption optimization",
			Tags:        []string{"shower-thoughts", "health"},
			Body: `What is the best way to tell if you should drink?
Drinking too much reduces productivity, but a lack of water reduces thinking ability.
A good start would be to drink when you feel *thirsty*.`,
		}
	)

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("creates HTML",
		func(t *testing.T) {
			buf := bytes.Buffer{}
			err := postRenderer.Render(&buf, Post)

			if err != nil {
				t.Fatal(err)
			}

			approvals.VerifyString(t, buf.String())
		})
	t.Run("make an index of posts",
		func(t *testing.T) {
			buf := bytes.Buffer{}
			posts := []blogrenderer.Post{
				{Title: "Post 1"},
				{Title: "Post 2"},
			}

			if err := postRenderer.RenderIndex(&buf, posts); err != nil {
				t.Fatal(err)
			}

			approvals.VerifyString(t, buf.String())
		})
}

func BenchmarkRenderer(b *testing.B) {
	var (
		Post = blogrenderer.Post{
			Title:       "Transformers",
			Description: "More than meets the eye",
			Tags:        []string{"robots", "in-disguise"},
			Body: `Autobots wage their battle to destroy
the evil forces of the Decepticons`,
		}
	)
	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, Post)
	}
}
