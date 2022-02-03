package utils

import (
	"context"
	"fmt"
	"io"
	"os"

	"os/exec"
	"strings"

	"github.com/docker/docker/client"
)

func SaveTarFromDockerImage(imageName, tarName string) error {
	dockerToTarCmd := fmt.Sprintf("docker save %s -o %s", imageName, tarName)
	exec.Command("/bin/sh", "-c", dockerToTarCmd).Output()
	return nil
}

func containsInSlice(slice []string, toSearch string) bool {
	for _, st := range slice {
		if strings.Contains(st, toSearch) {
			return true
		}
	}
	return false
}

func LoadTarToDocker(modifiedTarName string) (string, error) {
	file, err := os.Open(modifiedTarName + ".tar")
	if err != nil {
		return "", err
	}
	defer file.Close()
	tarReader := io.Reader(file)
	dockerClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}
	response, err := dockerClient.ImageLoad(context.Background(), tarReader, false)
	if err != nil {
		return "", err
	}
	outString := make([]byte, 1024)
	response.Body.Read(outString)
	response.Body.Close()
	return string(outString), nil
}
