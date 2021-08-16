package docker

import (
	"github.com/docker/docker/api/types"
)

func ExecContainer(containerID string) types.HijackedResponse {
	config := types.ExecConfig{
		Cmd:          []string{"bash"},
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Tty:          true,
		Detach:       false,
	}
	res, err := Client.ContainerExecCreate(ctx, containerID, config)
	if err != nil {
		println("Erreur1")
		panic(err)
	}

	res2, err := Client.ContainerExecAttach(ctx, res.ID, types.ExecStartCheck{
		Detach: false,
		Tty:    false,
	})
	if err != nil {
		println("Erreur3")
		panic(err)
	}

	return res2
}
