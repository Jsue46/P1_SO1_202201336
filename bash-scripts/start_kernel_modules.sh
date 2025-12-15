#!/bin/bash

# start_kernel_modules.sh
# Carga los módulos del kernel de tu proyecto

MOD_DIR="$(dirname "$0")/../modulo-kernel"

cd "$MOD_DIR" || exit 1

echo "[KERNEL] Cargando módulos..."

sudo insmod cont_info.ko
sudo insmod sys_info.ko

echo "[KERNEL] Módulos cargados correctamente."
