package blogposts

import (
	"bufio"
	"io"
	"io/fs"
)

type Post struct {
	Title, Description string
}

func NewBlogPost(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, f string) (Post, error) {
	postFile, err := fileSystem.Open(f)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	postTitle := readLine()[7:]
	postDescription := readLine()[13:]

	return Post{Title: postTitle, Description: postDescription}, nil
}
