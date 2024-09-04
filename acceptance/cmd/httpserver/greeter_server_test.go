package main_test

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	go_specs_greet "github.com/thomasonzhou/go-tdd/acceptance"
	"github.com/thomasonzhou/go-tdd/acceptance/specifications"
)

func TestGreeterServer(t *testing.T) {

	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../.", //acceptance folder
			Dockerfile:    "./cmd/httpserver/Dockerfile",
			PrintBuildLog: true,
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})

	driver := go_specs_greet.Driver{BaseUrl: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
