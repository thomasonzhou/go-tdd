package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	mdParser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: mdParser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newViewModel(p, r))
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.ReplaceAll(p.Title, " ", "-"))
}

func newViewModel(p Post, r *PostRenderer) postViewModel {
	return postViewModel{
		Post:     p,
		HTMLBody: template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil)),
	}
}
