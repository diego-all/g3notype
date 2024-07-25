import sys
import json

def main(json_path):
    with open(json_path, 'r') as archivo:
        datos_lista = json.load(archivo)
        datos = datos_lista[0]  # Primer elemento del array JSON
        tipo = datos['tipo']
        atributos = datos['atributos']

        matriz_atributos = []
        for nombre, detalle in atributos.items():
            tipo_dato = detalle['tipoDato']
            matriz_atributos.append([nombre, tipo_dato])

        return tipo, matriz_atributos

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Uso: python3 extractor/readMap.py <ruta_del_json>")
        sys.exit(1)
    json_path = sys.argv[1]
    tipo, matriz_atributos = main(json_path)
    
    # Salida como string y lista de listas
    print(tipo)
    for atributo in matriz_atributos:
        print(f"{atributo[0]}|{atributo[1]}")

