# extractor/readMap.py
import json
import sys

def main(json_path):
    with open(json_path, 'r') as archivo:
        datos_lista = json.load(archivo)
        datos = datos_lista[0] #Primer elemento del array JSON
        tipo = datos['tipo']
        print(f'Tipo: {tipo}')
        atributos = datos['atributos']
        for nombre, detalle in atributos.items():
            tipo_dato = detalle['tipoDato']
            print(f'Atributo: {nombre}, Tipo de dato: {tipo_dato}')

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Uso: python3 extractor/readMap.py <ruta_del_json>")
        sys.exit(1)
    json_path = sys.argv[1]
    main(json_path)
