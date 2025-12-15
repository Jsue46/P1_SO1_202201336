// sysinfo_reader.go
// Lee /proc/sysinfo_so1_202201336 y guarda métricas globales del sistema en SQLite

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// Estructura EXACTA según el JSON del módulo sys_info.c
type SysInfo struct {
	RamTotalKB uint64 `json:"ram_total_kb"`
	RamLibreKB uint64 `json:"ram_libre_kb"`
	RamUsadaKB uint64 `json:"ram_usada_kb"`
}

// Lee /proc y guarda en SQLite
func ReadAndStoreSystemInfo(db *sql.DB) {
	// 1. Leer el archivo /proc
	data, err := os.ReadFile("/proc/sysinfo_so1_202201336")
	if err != nil {
		log.Println("Error leyendo /proc/sysinfo:", err)
		return
	}

	// 2. Parsear JSON
	var info SysInfo
	err = json.Unmarshal(data, &info)
	if err != nil {
		log.Println("Error parseando JSON sysinfo:", err)
		return
	}

	// 3. Insertar en SQLite
	_, err = db.Exec(`
		INSERT INTO system_metrics (total_ram_kb, free_ram_kb, used_ram_kb)
		VALUES (?, ?, ?)
	`,
		info.RamTotalKB,
		info.RamLibreKB,
		info.RamUsadaKB,
	)

	if err != nil {
		log.Println("Error insertando system_metrics:", err)
		return
	}

	fmt.Printf(
		"[SYSINFO] Total: %d KB | Libre: %d KB | Usada: %d KB | %s\n",
		info.RamTotalKB,
		info.RamLibreKB,
		info.RamUsadaKB,
		time.Now().Format(time.RFC3339),
	)
}
