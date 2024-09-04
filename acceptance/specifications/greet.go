package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(string) (string, error)
}

func GreetSpecification(t *testing.T, greeter Greeter) {
	got, err := greeter.Greet("Tom")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello Tom")
}
