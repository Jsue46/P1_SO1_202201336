# MANUAL TÉCNICO  
**Proyecto 1 – Sistemas Operativos 1**

---

## 1. Objetivo del proyecto

El presente proyecto implementa un sistema de monitoreo de recursos a nivel de sistema operativo, integrando módulos del kernel de Linux, un daemon en espacio de usuario y una plataforma de visualización. El sistema permite obtener métricas globales del sistema y métricas por proceso asociadas a contenedores Docker, persistiendo la información para su análisis histórico.

---

## 2. Arquitectura general del sistema

El sistema está compuesto por cuatro capas principales:

1. **Kernel Linux**  
   Módulos personalizados exponen información del sistema y de procesos mediante el sistema de archivos `/proc`.

2. **Daemon en Go (User Space)**  
   Consume la información expuesta por el kernel, detecta contenedores Docker, calcula métricas dinámicas y persiste los datos.

3. **Base de datos SQLite**  
   Almacena métricas del sistema, métricas por contenedor y eventos generados por umbrales.

4. **Grafana**  
   Visualiza la información almacenada en SQLite mediante dashboards interactivos.



---

## 3. Módulos del kernel

### 3.1 Módulo `sys_info.c`

Este módulo expone información global del sistema a través del archivo:

