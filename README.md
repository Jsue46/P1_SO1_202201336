# Linux Kernel Resource Monitoring and Docker Process Management

Academic project developed for the **Operating Systems 1** course at the Universidad de San Carlos de Guatemala (Second Semester 2025).

This repository contains a hybrid resource monitoring system integrating custom Linux Kernel modules (kernel space) with a Go monitoring daemon (user space), SQLite persistence, and Grafana visualization.

---

## Overview

The system collects, aggregates, and persists low-level operating system metrics and container statistics without relying on standard third-party agents. Kernel modules expose real-time metrics via the `/proc` filesystem, which a Go daemon consumes periodically to calculate container CPU/RAM usage, detect anomalies, execute threshold-based policies, and stream records to SQLite.

---

## Architecture

```
[ Linux Kernel Space ]
  ├── sys_info.c     -> Exposes system-wide RAM metrics via /proc/sys_info
  └── cont_info.c    -> Exposes process/cgroup memory metrics via /proc/cont_info
           │
           │ Read metrics (/proc)
           v
[ Go Monitoring Daemon (User Space) ]
  ├── Cgroup parser for container identification
  ├── Real-time CPU and RAM differential calculator
  ├── Automated threshold policy engine (WARN, STOP, KILL)
  └── SQLite database writer (WAL mode)
           │
           v
[ SQLite Database (WAL Mode) ]
           │
           v Read metrics
[ Grafana Dashboard (Local Container) ]
```

---

## Key Features

- Custom C kernel modules integrated directly with the Linux `/proc` filesystem
- Low-overhead container identification and metrics extraction via cgroups
- Differential CPU utilization algorithms based on system timer ticks
- Automatic enforcement actions on runaway containers exceeding RAM/CPU thresholds
- Concurrent persistence with SQLite in WAL (Write-Ahead Logging) mode
- Real-time Grafana dashboards displaying historical memory trends, top processes, and container lifecycle events

---

## Technologies

- **Kernel Programming:** C, Linux Kernel Headers
- **User Space Daemon:** Go (Golang)
- **Database:** SQLite (WAL mode)
- **Visualization:** Grafana
- **Containerization:** Docker, Docker Compose

---

## Repository Structure

```
P1_SO1_202201336/
├── modules/                  # C kernel module source code
│   ├── sys_info.c
│   └── cont_info.c
├── daemon/                   # Go monitoring daemon source code
├── dashboards/               # Grafana dashboard templates
├── docs/                     # Technical and user manuals
├── docker-compose.yaml       # Grafana service orchestration
└── README.md
```

---

## Academic Context

| Field       | Detail                                         |
|-------------|------------------------------------------------|
| Course      | Operating Systems 1 (SO1)                      |
| Institution | Universidad de San Carlos de Guatemala         |
| Semester    | Second Semester 2025                           |
| Author      | Estuardo Josue Vaquiax Reyes (202201336)       |
