#!/bin/bash

# cleanup_container.sh
# Elimina un contenedor por su ID

CONTAINER_ID="$1"

if [ -z "$CONTAINER_ID" ]; then
    echo "Error: no se proporcionó un container ID"
    exit 1
fi

echo "[CLEANUP] Deteniendo contenedor $CONTAINER_ID..."
docker stop "$CONTAINER_ID" >/dev/null 2>&1

echo "[CLEANUP] Eliminando contenedor $CONTAINER_ID..."
docker rm "$CONTAINER_ID" >/dev/null 2>&1

echo "[CLEANUP] Contenedor $CONTAINER_ID detenido y eliminado."
exit 0
