package main

import (
	svg "github.com/thomasonzhou/go-tdd/math/svg"
	"os"
	"time"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}
