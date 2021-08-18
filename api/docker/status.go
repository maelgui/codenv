package docker

import "github.com/docker/docker/api/types"

func StartContainer(containerID string) error {
	return Client.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
}

func RestartContainer(containerID string) error {
	return Client.ContainerRestart(ctx, containerID, nil)
}

func StopContainer(containerID string) error {
	return Client.ContainerStop(ctx, containerID, nil)
}
