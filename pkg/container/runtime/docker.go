/*
Copyright 2022 k0s authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package runtime

import (
	"fmt"
	"os/exec"
	"strings"
)

var _ ContainerRuntime = (*DockerRuntime)(nil)

type DockerRuntime struct {
	criSocketPath string
}

func (d *DockerRuntime) ListContainers() ([]string, error) {
	out, err := exec.Command("docker", "--host", d.criSocketPath, "ps", "-a", "--filter", "name=k8s_", "-q").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: output: %s: %w", string(out), err)
	}
	return strings.Fields(string(out)), nil
}

func (d *DockerRuntime) RemoveContainer(id string) error {
	out, err := exec.Command("docker", "--host", d.criSocketPath, "rm", "--volumes", id).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to remove container %s: output: %s: %w", id, string(out), err)
	}
	return nil
}

func (d *DockerRuntime) StopContainer(id string) error {
	out, err := exec.Command("docker", "--host", d.criSocketPath, "stop", id).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stop running container %s: output: %s: %w", id, string(out), err)
	}
	return nil
}
