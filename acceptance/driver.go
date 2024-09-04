package go_specs_greet

import (
	"fmt"
	"io"
	"net/http"
)

type Driver struct {
	BaseUrl string
}

func (d Driver) Greet() (string, error) {
	res, err := http.Get(d.BaseUrl + "/greet")
	if err != nil {
		return "", fmt.Errorf("failed to get greeting %w", err)
	}
	defer res.Body.Close()
	greeting, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read greeting body %w", err)
	}
	return string(greeting), nil
}
