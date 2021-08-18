package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

func CreateContainer(image string) string {
	_, err := Client.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	endpointsConfig := map[string]*network.EndpointSettings{
		"codenv_network": &network.EndpointSettings{},
	}

	resp, err := Client.ContainerCreate(ctx, &container.Config{
		Image: image,
		Cmd:   []string{"echo", "hello world"},
		Volumes: map[string]struct{}{
			"/home": struct{}{},
		},
	}, nil, &network.NetworkingConfig{EndpointsConfig: endpointsConfig}, nil, "")
	if err != nil {
		panic(err)
	}

	return resp.ID
}

func DeleteContainer(containerID string) {
	options := types.ContainerRemoveOptions{
		RemoveVolumes: true,
		RemoveLinks:   false,
		Force:         true,
	}

	err := Client.ContainerRemove(ctx, containerID, options)
	if err != nil {
		panic(err)
	}
}

func RetrieveContainer(id string) types.ContainerJSON {
	res, err := Client.ContainerInspect(ctx, id)
	if err != nil {
		panic(err)
	}

	return res
}
