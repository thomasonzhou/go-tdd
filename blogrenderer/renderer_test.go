package blogrenderer_test

import (
	"bytes"
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
A good start would be to drink when you feel thirsty.`,
		}
	)

	t.Run("creates HTML",
		func(t *testing.T) {
			buf := bytes.Buffer{}
			err := blogrenderer.Render(&buf, Post)

			if err != nil {
				t.Fatal(err)
			}

			approvals.VerifyString(t, buf.String())
		})
}
