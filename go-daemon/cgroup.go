package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)


// DetectDockerContainerID detecta si un PID pertenece a un contenedor Docker
// usando cgroup v2 con systemd (docker-<id>.scope)
func DetectDockerContainerID(pid int) (string, error) {
	path := "/proc/" + strconv.Itoa(pid) + "/cgroup"

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Soportar múltiples formatos de cgroup (v1 y v2) comunes:
	// v2 systemd:   0::/system.slice/docker-<ID>.scope
	// v1 docker:    1:cpuacct:/docker/<ID>
	// v1 systemd:   1:name=systemd:/system.slice/docker-<ID>.scope
	// containerd:   0::/system.slice/docker-<ID>.scope o /docker/<ID>
	// crio/kube:    .../kubepods/.../pod.../containerd-<ID>.scope (no Docker)

	lines := strings.Split(string(data), "\n")

	// Regex para extraer IDs en varios formatos
	// docker-<ID>.scope
	reScope := regexp.MustCompile(`docker-([a-f0-9]{12,64})\.scope`)
	// /docker/<ID>
	reDocker := regexp.MustCompile(`/docker/([a-f0-9]{12,64})`)
	// containerd-<ID>.scope (por si viene con containerd pero usando docker runtime)
	reContainerdScope := regexp.MustCompile(`containerd-([a-f0-9]{12,64})\.scope`)

	for _, line := range lines {
		if line == "" {
			continue
		}

		if m := reScope.FindStringSubmatch(line); len(m) == 2 {
			return m[1], nil
		}
		if m := reDocker.FindStringSubmatch(line); len(m) == 2 {
			return m[1], nil
		}
		if m := reContainerdScope.FindStringSubmatch(line); len(m) == 2 {
			return m[1], nil
		}
	}

	return "", nil
}
