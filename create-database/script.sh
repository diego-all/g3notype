# Validar y eliminar el archivo create_database.sql si existe
if [ -f create_database.sql ]; then
    rm create_database.sql
fi

cp ../projectTest/database/up.sql create_database.sql

# Validar y eliminar el directorio de destino si existe
if [ -d ~/PROBAR-GENERADA ]; then
    rm -rf ~/PROBAR-GENERADA
fi

cp -R ../projectTest ~/PROBAR-GENERADA

sqlite3 data.sqlite < create_database.sql

# Validar y eliminar el archivo data.sqlite en el destino si existe
if [ -f ~/PROBAR-GENERADA/projectTest/data.sqlite ]; then
    rm ~/PROBAR-GENERADA/projectTest/data.sqlite
fi

cp data.sqlite ~/PROBAR-GENERADA/projectTest
