# P1_SO1_202201336# Proyecto 1 – Sistemas Operativos 1  
## Monitoreo y Gestión Inteligente de Procesos y Contenedores en Linux

Universidad de San Carlos de Guatemala  
Facultad de Ingeniería  
Ingeniería en Ciencias y Sistemas  

---

## 📌 Descripción General

Este proyecto implementa un sistema integral de monitoreo, análisis y gestión automática de procesos del sistema operativo y contenedores Docker en Linux.  
La solución combina programación a bajo nivel mediante módulos de kernel en C y programación a alto nivel utilizando un daemon desarrollado en Go, con visualización de métricas a través de Grafana.

La información es recolectada directamente desde el kernel, expuesta mediante el sistema de archivos `/proc`, procesada en espacio de usuario y utilizada para la toma de decisiones automáticas que garantizan la estabilidad del sistema.

---

## 🎯 Objetivo General

Desarrollar un sistema que permita monitorear procesos del sistema y procesos asociados a contenedores Docker, analizar su consumo de recursos y ejecutar acciones automáticas para mantener la estabilidad del sistema.

---

## 🎯 Objetivos Específicos

- Desarrollar módulos de kernel funcionales en C.
- Extraer métricas de procesos directamente desde el kernel.
- Exponer información mediante archivos virtuales en `/proc`.
- Desarrollar un daemon en Go para el análisis de métricas.
- Implementar decisiones automáticas para la gestión de contenedores.
- Persistir información histórica utilizando SQLite.
- Visualizar métricas mediante dashboards en Grafana.
- Automatizar la generación de contenedores mediante cronjobs.

---

## 🧠 Arquitectura del Sistema

El sistema está compuesto por los siguientes componentes:

1. **Módulos de Kernel**
   - Módulo de procesos de contenedores
   - Módulo de procesos del sistema

2. **Daemon en Go**
   - Lectura y parseo de archivos `/proc`
   - Análisis de métricas
   - Gestión automática de contenedores
   - Registro de eventos en SQLite

3. **Cronjob**
   - Generación automática de contenedores cada minuto

4. **Grafana**
   - Visualización de métricas históricas y en tiempo real

---

## 📁 Estructura del Repositorio

```text
proyecto-1/
├── modulo-kernel/
│   ├── continfo/
│   │   ├── continfo.c
│   │   └── Makefile
│   ├── sysinfo/
│   │   ├── sysinfo.c
│   │   └── Makefile
│
├── go-daemon/
│   ├── main.go
│   ├── parser.go
│   ├── analyzer.go
│   ├── docker_manager.go
│   ├── database.go
│   └── cron.go
│
├── bash/
│   ├── load_modules.sh
│   ├── unload_modules.sh
│   └── generate_containers.sh
│
├── dashboard/
│   ├── docker-compose.yml
│   └── grafana/
│
└── documentacion/
    ├── manual_tecnico.md
    └── manual_usuario.md
