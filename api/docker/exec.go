package docker

import (
	"github.com/docker/docker/api/types"
)

func OpenTerminal(containerID string) string {
	config := types.ExecConfig{
		Cmd:          []string{"bash"},
		AttachStdin:  true,
		AttachStderr: true,
		AttachStdout: true,
		Tty:          true,
		Detach:       false,
		Env: []string{"TERM=xterm-256color"},
	}
	res, err := Client.ContainerExecCreate(ctx, containerID, config)
	if err != nil {
		println("Erreur1")
		panic(err)
	}

	return res.ID
}

func AttachExec(execID string) types.HijackedResponse {

	res2, err := Client.ContainerExecAttach(ctx, execID, types.ExecStartCheck{
		Detach: false,
		Tty:    false,
	})
	if err != nil {
		println("Erreur3")
		panic(err)
	}

	return res2
}
