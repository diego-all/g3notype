#!/bin/bash

# Validar si sqlite3 está instalado, de lo contrario, instalarlo.
if ! command -v sqlite3 &> /dev/null
then
    echo "sqlite3 no está instalado. Instalando sqlite3..."
    
    # Instalar sqlite3 según el sistema operativo
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        sudo apt-get update
        sudo apt-get install -y sqlite3
    elif [[ "$OSTYPE" == "darwin"* ]]; then
        brew install sqlite3
    else
        echo "Sistema operativo no soportado para instalación automática."
        exit 1
    fi
    
    # Verificar si la instalación fue exitosa
    if command -v sqlite3 &> /dev/null
    then
        echo "sqlite3 se ha instalado correctamente."
    else
        echo "Error: sqlite3 no se pudo instalar."
        exit 1
    fi
else
    echo "sqlite3 ya está instalado."
fi

# Crear la base de datos en la carpeta database
temp_db_path="./data.sqlite"

if sqlite3 "$temp_db_path" < up.sql; then
    echo "Base de datos creada exitosamente en la carpeta database."

    # Mover la base de datos a la raíz del directorio del proyecto
    mv "$temp_db_path" ../data.sqlite
    echo "Base de datos movida exitosamente a la raíz del directorio."
else
    echo "Error: No se pudo crear la base de datos."
    exit 1
fi
