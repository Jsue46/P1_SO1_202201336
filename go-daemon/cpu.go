// cpu.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPUStat struct {
	ProcTime  uint64
	TotalTime uint64
}

func getProcessCPU(pid int) (uint64, error) {
	statPath := fmt.Sprintf("/proc/%d/stat", pid)

	data, err := os.ReadFile(statPath)
	if err != nil {
		return 0, err
	}

	fields := strings.Fields(string(data))
	if len(fields) < 17 {
		return 0, fmt.Errorf("stat corrupto para pid %d", pid)
	}

	utime, _ := strconv.ParseUint(fields[13], 10, 64)
	stime, _ := strconv.ParseUint(fields[14], 10, 64)

	return utime + stime, nil
}

func getTotalCPU() (uint64, error) {
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		return 0, err
	}

	line := strings.Split(string(data), "\n")[0]
	fields := strings.Fields(line)

	var total uint64
	for i := 1; i < len(fields); i++ {
		v, _ := strconv.ParseUint(fields[i], 10, 64)
		total += v
	}

	return total, nil
}

func CaptureSnapshot(pid int) (CPUStat, error) {
	proc, err := getProcessCPU(pid)
	if err != nil {
		return CPUStat{}, err
	}

	total, err := getTotalCPU()
	if err != nil {
		return CPUStat{}, err
	}

	return CPUStat{
		ProcTime:  proc,
		TotalTime: total,
	}, nil
}

func CalculateCPUPercent(pid int, prev, curr CPUStat) float64 {
	deltaProc := float64(curr.ProcTime - prev.ProcTime)
	deltaTotal := float64(curr.TotalTime - prev.TotalTime)

	if deltaTotal <= 0 {
		return 0
	}

	return (deltaProc / deltaTotal) * 100
}
