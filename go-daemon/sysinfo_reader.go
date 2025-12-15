// sysinfo_reader.go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type SysInfo struct {
	RamTotalKB uint64 `json:"ram_total_kb"`
	RamLibreKB uint64 `json:"ram_libre_kb"`
	RamUsadaKB uint64 `json:"ram_usada_kb"`
}

func ReadAndStoreSystemInfo() {
	data, err := os.ReadFile("/proc/sysinfo_so1_202201336")
	if err != nil {
		fmt.Println("Error leyendo /proc/sysinfo:", err)
		return
	}

	var info SysInfo
	if err := json.Unmarshal(data, &info); err != nil {
		fmt.Println("Error parseando JSON sysinfo:", err)
		return
	}

	InsertSystemMetrics(info.RamTotalKB, info.RamLibreKB, info.RamUsadaKB)
}
