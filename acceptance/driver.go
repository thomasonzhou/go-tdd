package go_specs_greet

import (
	"fmt"
	"io"
	"net/http"
)

type Driver struct {
	BaseUrl string
	Client  *http.Client
}

func (d Driver) Greet(name string) (string, error) {
	res, err := d.Client.Get(d.BaseUrl + "/greet?name=" + name)
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
