package docker

import "github.com/docker/docker/api/types"

func StartContainer(containerID string) {
	if err := Client.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}

func RestartContainer(containerID string) {
	if err := Client.ContainerRestart(ctx, containerID, nil); err != nil {
		panic(err)
	}
}

func StopContainer(containerID string) {
	if err := Client.ContainerStop(ctx, containerID, nil); err != nil {
		panic(err)
	}
}
