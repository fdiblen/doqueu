package model

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// Container
type Container struct {
	ID     string `json:"id" example:"25fe01fcafc5"`
	Image  string `json:"image" example:"alpine"`
	Status string `json:"status" example:"Running"`
}

// ContainersAll
func ContainersAll() ([]Container, error) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	var ContainerList []Container
	for _, _container := range containers {
		ContainerList = append(ContainerList, Container{_container.ID, _container.Image, _container.Status})
	}

	return ContainerList, nil
}

// ContainerOne
func ContainerOne(id string) (*Container, error) {
	containers, err := ContainersAll()

	for _, v := range containers {
		if id == v.ID {
			return &v, nil
		}
	}
	return nil, err
}

// ContainerStart
func ContainerStart(imagename string, command string) (string, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imagename,
		Cmd:   strings.Fields(command),
	}, nil, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, Follow: true})
	if err != nil {
		return "", err
	}

	buffer, err := io.ReadAll(out)
	if err != nil && err != io.EOF {
		return "", err
	}

	fmt.Println("Container output: ", string(buffer))

	return string(buffer), nil
}

// ContainerEnd
func ContainerEnd(id string) (string, error) {

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	fmt.Print("Stopping container ", id, "... ")
	if err := cli.ContainerStop(ctx, id, nil); err != nil {
		panic(err)
	}

	return "Success", nil
}
