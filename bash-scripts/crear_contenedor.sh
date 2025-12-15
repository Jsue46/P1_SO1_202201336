#!/bin/bash

# Nombre único usando timestamp
NAME="ct_$(date +%s)"

# Crear contenedor liviano que consuma algo de RAM
docker run -d --name $NAME alpine sh -c "sleep 300"
