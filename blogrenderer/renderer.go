package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(buf io.Writer, p Post) error {
	var tagItems string
	for _, tag := range p.Tags {
		tagItems += fmt.Sprintf("<li>%v</li>", tag)
	}

	_, err := fmt.Fprintf(buf, `<h1>%v</h1>
<p>%v</p>
<ul>%v</ul>`,
		p.Title, p.Description, tagItems)
	return err
}
