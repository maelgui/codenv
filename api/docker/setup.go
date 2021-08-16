package docker

import (
	"context"

	"github.com/docker/docker/client"
)

var ctx context.Context
var Client *client.Client

func ConnectDocker() {
	ctx = context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	Client = cli
}
