package main

import (
	"log"
	"os"

	blogposts "github.com/thomasonzhou/go-tdd/blogposts"
)

func main() {
	posts, err := blogposts.NewBlogPost(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
