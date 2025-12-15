package main

import (
	"fmt"
	"os/exec"
)

// Detiene un contenedor Docker
func StopContainer(containerID string) {
	fmt.Println("🟡 Deteniendo contenedor:", containerID)

	cmd := exec.Command("docker", "stop", containerID)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error al detener contenedor:", err)
	}
}

// Elimina un contenedor Docker
func KillContainer(containerID string) {
	fmt.Println("🔴 Eliminando contenedor:", containerID)

	cmd := exec.Command("docker", "rm", "-f", containerID)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error al eliminar contenedor:", err)
	}
}
