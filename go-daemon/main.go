package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Daemon SO1 iniciado...")

	InitDB()

	for {
		// 1) Guardar métricas globales del sistema (Total/Libre/Usada)
		ReadAndStoreSystemInfo()

		kd, err := ReadKernelProcesses()
		if err != nil {
			fmt.Println("Error leyendo kernel:", err)
			time.Sleep(3 * time.Second)
			continue
		}

		fmt.Printf("\nProcesos detectados: %d\n", kd.TotalProcesos)

		debugPrinted := false

		for _, p := range kd.Procesos {
			containerID, _ := DetectDockerContainerID(p.PID)
			if containerID == "" {
				if !debugPrinted {
					path := fmt.Sprintf("/proc/%d/cgroup", p.PID)
					data, err := os.ReadFile(path)
					if err != nil {
						fmt.Printf("PID %d sin cgroup: %v\n", p.PID, err)
					} else {
						fmt.Printf("PID %d cgroup:\n%s\n", p.PID, string(data))
					}
					debugPrinted = true
				}
				continue
			}

			snap1, err := CaptureSnapshot(p.PID)
			if err != nil {
				continue
			}

			time.Sleep(1 * time.Second)

			snap2, err := CaptureSnapshot(p.PID)
			if err != nil {
				continue
			}

			cpu := CalculateCPUPercent(p.PID, snap1, snap2)

			fmt.Printf(" → PID %d (%s) en contenedor %s\n", p.PID, p.Nombre, containerID)
			fmt.Printf("      RSS: %d KB  CPU%%: %.2f%%  RAM%%: %d%%\n", p.RSS, cpu, p.RAMPercent)

			InsertContainerMetrics(
				containerID,
				p.PID,
				p.RSS,
				p.VSZ,
				p.RAMPercent, // <-- YA NO 0
				cpu,
			)

			ApplyThresholdActions(containerID, p.RAMPercent, cpu)
		}

		fmt.Println("Ciclo completado. Esperando 5s...")
		time.Sleep(5 * time.Second)
	}
}
