// database.go
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Ruta al SQLite
const dbPath = "../docker/data/metrics.db"

func InitDB() {
	var err error

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	// Configuración CLAVE para SQLite + Grafana (concurrencia)
	db.Exec("PRAGMA journal_mode=WAL;")
	db.Exec("PRAGMA synchronous=NORMAL;")
	db.Exec("PRAGMA busy_timeout=5000;")

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// ===== Tablas =====

	createSystemMetrics := `
	CREATE TABLE IF NOT EXISTS system_metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		total_ram_kb INTEGER,
		free_ram_kb  INTEGER,
		used_ram_kb  INTEGER,
		timestamp    DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	createMetrics := `
	CREATE TABLE IF NOT EXISTS metrics (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		container_id TEXT,
		pid          INTEGER,
		rss_kb       INTEGER,
		vsz_kb       INTEGER,
		ram_percent  REAL,
		cpu_percent  REAL,
		timestamp    DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	createEvents := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		container_id TEXT,
		metric       TEXT,   -- "CPU" o "RAM"
		action       TEXT,   -- "WARN" | "STOP" | "KILL"
		timestamp    DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	// ===== Índices (ayudan a Grafana y consultas Top 5) =====
	createIndexes := []string{
		`CREATE INDEX IF NOT EXISTS idx_system_ts ON system_metrics(timestamp);`,
		`CREATE INDEX IF NOT EXISTS idx_metrics_ts ON metrics(timestamp);`,
		`CREATE INDEX IF NOT EXISTS idx_metrics_container ON metrics(container_id);`,
		`CREATE INDEX IF NOT EXISTS idx_events_ts ON events(timestamp);`,
		`CREATE INDEX IF NOT EXISTS idx_events_action ON events(action);`,
	}

	if _, err = db.Exec(createSystemMetrics); err != nil {
		log.Fatal("Error creando system_metrics:", err)
	}
	if _, err = db.Exec(createMetrics); err != nil {
		log.Fatal("Error creando metrics:", err)
	}
	if _, err = db.Exec(createEvents); err != nil {
		log.Fatal("Error creando events:", err)
	}
	for _, q := range createIndexes {
		_, _ = db.Exec(q)
	}

	fmt.Println("Base de datos inicializada correctamente.")
}

// Inserta métricas globales del sistema (RAM total/libre/usada)
func InsertSystemMetrics(totalKB, freeKB, usedKB uint64) {
	if db == nil {
		return
	}
	_, err := db.Exec(`
		INSERT INTO system_metrics (total_ram_kb, free_ram_kb, used_ram_kb)
		VALUES (?, ?, ?)
	`, totalKB, freeKB, usedKB)

	if err != nil {
		fmt.Println("Error insertando system_metrics:", err)
	}
}

// Inserta métricas de contenedor (PID + ID + RSS/VSZ + CPU/RAM%)
func InsertContainerMetrics(
	containerID string,
	pid int,
	rssKB int,
	vszKB int,
	ramPercent int,
	cpuPercent float64,
) {
	if db == nil {
		return
	}

	_, err := db.Exec(`
		INSERT INTO metrics (container_id, pid, rss_kb, vsz_kb, ram_percent, cpu_percent)
		VALUES (?, ?, ?, ?, ?, ?)
	`, containerID, pid, rssKB, vszKB, ramPercent, cpuPercent)

	if err != nil {
		fmt.Println("Error insertando métricas:", err)
	}
}

// Inserta eventos del sistema (WARN/STOP/KILL) para CPU o RAM
func InsertEvent(containerID, metric, action string) {
	if db == nil {
		return
	}

	_, err := db.Exec(`
		INSERT INTO events (container_id, metric, action)
		VALUES (?, ?, ?)
	`, containerID, metric, action)

	if err != nil {
		fmt.Println("Error insertando evento:", err)
	}
}
