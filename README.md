# Proyecto 1 – Sistemas Operativos 1  
**Monitoreo de Recursos del Sistema y Contenedores Docker**

---

## 📌 Descripción general

Este proyecto implementa un sistema de monitoreo de recursos a nivel de sistema operativo, integrando componentes en **kernel-space** y **user-space**. El sistema permite recolectar métricas del sistema y de procesos asociados a contenedores Docker, almacenarlas de forma persistente y visualizarlas mediante dashboards en Grafana.

La solución combina módulos del kernel de Linux, un daemon desarrollado en Go, una base de datos SQLite y Grafana como herramienta de visualización.

---

## 🎯 Objetivo del proyecto

- Obtener métricas globales del sistema (memoria RAM).
- Monitorear procesos asociados a contenedores Docker.
- Calcular uso de CPU y RAM por contenedor.
- Persistir información histórica para análisis.
- Visualizar métricas y eventos en tiempo real.
- Aplicar acciones automáticas sobre contenedores que excedan umbrales definidos.

---

## 🧱 Arquitectura del sistema

El sistema está compuesto por cuatro capas claramente diferenciadas:

1. **Módulos del Kernel (C)**  
   Exponen información del sistema y de procesos mediante archivos en `/proc`.

2. **Daemon en Go (User Space)**  
   Consume la información del kernel, detecta contenedores Docker, calcula métricas dinámicas y persiste los datos.

3. **Base de Datos SQLite**  
   Almacena métricas del sistema, métricas por contenedor y eventos generados por el sistema.

4. **Grafana**  
   Consulta la base de datos y presenta dashboards interactivos para el análisis visual.

<!-- IMAGEN: Diagrama de arquitectura general del sistema -->

---

## 🧩 Componentes principales

### 🔹 Módulos del kernel
- **sys_info.c**  
  Expone métricas globales de memoria RAM del sistema.

- **cont_info.c**  
  Expone información de procesos de usuario (PID, RSS, VSZ y porcentaje de RAM).

### 🔹 Daemon en Go
- Lee información desde `/proc`.
- Detecta contenedores Docker mediante `cgroup`.
- Calcula uso de CPU por proceso.
- Aplica umbrales y acciones automáticas (WARN, STOP, KILL).
- Inserta métricas en SQLite.

### 🔹 Persistencia y visualización
- **SQLite** con modo WAL para concurrencia.
- **Grafana** como plataforma de visualización en modo solo lectura.

---

## 📊 Métricas visualizadas

El dashboard principal presenta las siguientes métricas:

- Memoria RAM total del sistema
- Memoria RAM libre
- Memoria RAM usada
- Uso de RAM a lo largo del tiempo
- Uso de CPU a lo largo del tiempo
- Top 5 contenedores con mayor uso de RAM
- Top 5 contenedores con mayor uso de CPU
- Contenedores eliminados a lo largo del tiempo

---

## 📘 Manual Técnico

El **Manual Técnico** describe en detalle la implementación interna del sistema y está dirigido a desarrolladores o evaluadores técnicos.

Incluye:

- Diseño de los módulos del kernel.
- Comunicación mediante `/proc`.
- Arquitectura del daemon en Go.
- Cálculo de métricas de CPU y RAM.
- Modelo de datos en SQLite.
- Manejo de concurrencia y consistencia.
- Integración con Grafana.
- Limitaciones y mejoras propuestas.


---

## 🚀 Ejecución general del proyecto

De forma resumida, el flujo de ejecución es el siguiente:

1. Compilación e inserción de los módulos del kernel.
2. Ejecución del daemon en Go.
3. Levantamiento de Grafana mediante Docker Compose.
4. Acceso al dashboard desde el navegador.

Los pasos detallados se encuentran en el **Manual de Instalación**.

---

## 👤 Manual de Usuario

El **Manual de Usuario** está orientado al uso del sistema desde el punto de vista del usuario final, sin necesidad de conocimientos técnicos avanzados.

Incluye:

- Acceso a Grafana.
- Descripción del dashboard.
- Interpretación de cada panel.
- Uso del rango de tiempo.
- Visualización de eventos y acciones automáticas.

📄 Archivo:  


---

## 🛠 Tecnologías utilizadas

- **Lenguaje C** – Desarrollo de módulos del kernel
- **Go** – Daemon de monitoreo
- **SQLite** – Persistencia de datos
- **Docker / Docker Compose** – Contenerización
- **Grafana** – Visualización de métricas
- **Linux /proc** – Comunicación kernel ↔ user-space

---

## ⚠️ Alcance y consideraciones

- El proyecto está diseñado con fines académicos.
- Las métricas de CPU se calculan mediante muestreo temporal.
- El porcentaje de RAM se calcula respecto al total del sistema.
- SQLite es adecuado para el alcance del proyecto, pero no para entornos productivos de alta carga.

---

## 📄 Documentación incluida

- `README.md`
- `MANUAL_INSTALACION.md`
- `MANUAL_USUARIO.md`
- `MANUAL_TECNICO.md`

---

## ✅ Conclusión

El proyecto demuestra la integración de conceptos fundamentales de Sistemas Operativos, como monitoreo de procesos, gestión de memoria, comunicación kernel–user space y concurrencia. La solución implementada es funcional, modular y cumple con los requerimientos establecidos en el enunciado del curso.

