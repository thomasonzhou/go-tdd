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
	_, err := fmt.Fprintf(buf, `<h1>%v</h1>
<p>%v</p>`,
		p.Title, p.Description)
	return err
}
