#!/bin/bash

# cron_generator.sh
# Genera contenedores de prueba aleatorios (alto o bajo consumo)

# Probabilidad 50/50
R=$(( RANDOM % 2 ))

if [ "$R" -eq 0 ]; then
    echo "[CRON] Lanzando contenedor BAJO consumo..."
    docker run -d --name test_bajo_$(date +%s) bajo_consumo
else
    echo "[CRON] Lanzando contenedor ALTO consumo..."
    docker run -d --name test_alto_$(date +%s) alto_consumo
fi

exit 0
