import json

# Asumiendo que el JSON está en un archivo llamado 'classes.json' en la carpeta 'inputs'
with open('inputs/classes.json', 'r') as archivo:
    # Carga el JSON como una lista
    datos_lista = json.load(archivo)

    # Accede al primer objeto de la lista (si solo hay uno)
    datos = datos_lista[0]

    # Extraer el tipo
    tipo = datos['tipo']
    print(f'Tipo: {tipo}')

    # Extraer los atributos y sus tipos de dato
    atributos = datos['atributos']
    for nombre, detalle in atributos.items():
        tipo_dato = detalle['tipoDato']
        print(f'Atributo: {nombre}, Tipo de dato: {tipo_dato}')