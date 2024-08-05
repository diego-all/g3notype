#!/bin/sh

# Obtener la ruta del script
script_dir=$(cd "$(dirname "$0")" && pwd)

# Validar si sqlite3 está instalado, de lo contrario, instalarlo.
if ! command -v sqlite3 > /dev/null 2>&1; then
    echo "sqlite3 no está instalado. Instalando sqlite3..."
    
    # Instalar sqlite3 según el sistema operativo
    if [ "$OSTYPE" = "linux-gnu" ]; then
        sudo apt-get update
        sudo apt-get install -y sqlite3
    elif [ "$OSTYPE" = "darwin" ]; then
        brew install sqlite3
    else
        echo "Sistema operativo no soportado para instalación automática."
        exit 1
    fi
    
    # Verificar si la instalación fue exitosa
    if command -v sqlite3 > /dev/null 2>&1; then
        echo "sqlite3 se ha instalado correctamente."
    else
        echo "Error: sqlite3 no se pudo instalar."
        exit 1
    fi
else
    echo "sqlite3 ya está instalado."
fi

# Crear la base de datos en la misma carpeta que el script
temp_db_path="$script_dir/data.sqlite"

if sqlite3 "$temp_db_path" < "$script_dir/up.sql"; then
    echo "Base de datos creada exitosamente en la carpeta /database."

    # Copiar la base de datos a la raíz del directorio del proyecto
    cp "$temp_db_path" "$script_dir/../data.sqlite"
    echo "Base de datos copiada exitosamente a la raíz del directorio del proyecto."
else
    echo "Error: No se pudo crear la base de datos."
    exit 1
fi
