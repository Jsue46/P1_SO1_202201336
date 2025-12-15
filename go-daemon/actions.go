// thresholds.go
package main

import "fmt"

func ApplyThresholdActions(containerID string, ramPercent int, cpuPercent float64) {

	// RAM
	if ramPercent >= cfg.RamKill {
		fmt.Println("🚨 RAM KILL:", containerID)
		InsertEvent(containerID, "RAM", "KILL")
		KillContainer(containerID)
		return
	}

	if ramPercent >= cfg.RamStop {
		fmt.Println("⚠️ RAM STOP:", containerID)
		InsertEvent(containerID, "RAM", "STOP")
		StopContainer(containerID)
		return
	}

	if ramPercent >= cfg.RamWarn {
		fmt.Println("ℹ️ RAM WARN:", containerID)
		InsertEvent(containerID, "RAM", "WARN")
	}

	// CPU
	if cpuPercent >= cfg.CPUKill {
		fmt.Println("🚨 CPU KILL:", containerID)
		InsertEvent(containerID, "CPU", "KILL")
		KillContainer(containerID)
		return
	}

	if cpuPercent >= cfg.CPUStop {
		fmt.Println("⚠️ CPU STOP:", containerID)
		InsertEvent(containerID, "CPU", "STOP")
		StopContainer(containerID)
		return
	}

	if cpuPercent >= cfg.CPUWarn {
		fmt.Println("ℹ️ CPU WARN:", containerID)
		InsertEvent(containerID, "CPU", "WARN")
	}
}
