# G3notype (run-from-gh)

## Para su ejecución se debe especificar:

init       (Iniciar la generación del scaffolding.)
--config  (Ruta al archivo de configuración JSON con el modelo)
--db sqlite  (Base de datos a utilizar)
projectTest  (Nombre del proyecto)

### Input JSON (--config)

    {
      "tipo": "Book",
      "atributos": {
        "nombre": {
          "tipoDato": "string"
        },
        "descripcion": {
          "tipoDato": "string"
        },
        "precio": {
          "tipoDato": "integer"
        },
        "cantidad": {
          "tipoDato": "integer"
        },
        "random": {
          "tipoDato": "string"
        }
      }
    },


### Local

    go run github.com/diego-all/run-from-gh@latest init

    go run github.com/diego-all/run-from-gh@latest init -h

    go run main.go init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest


### From remote repository

    go run github.com/diego-all/run-from-gh@v0.1.1 init --db sqlite --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest


## Extractor (Python)

  Needs python > 3.7

  Uso: python3 extractor/readMap.py <ruta_del_json>

  Uso: python3 extractor/readMap.py /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json


## Test generated API

    cp -R projectTest /home/diegoall/PROBAR-GENERADA


## Analizar el tema del repositorio

    go run ./cmd/api
    go build ./cmd/api
    go build -o productsAPI ./cmd/api

    cmd/api/handlers-Book.go:8:2: no required module provides package github.com/diego-all/books-API/internal; to add it:

    Quiza deba cambirse el nombre para poder probar y no busque modulo.

    Funciona sin repositorio. Analizar como hacer que funcione con repositorio y crearlo con git. (analizar como crea el repo git init y demas)


## Ajuste descriptivo (book.Name vs book.nombre) crear tipos (Nombre)


## Se pierden las identaciones o tabulaciones en algunos casos (Ya ajustados)
Funciona bien, al parecer vscode organiza el script.


## Llenar los inserts

  Se hizo con Gemini (enpalmar con el generador o solicitar al usuario si desea crearlos)


## 


## 

lsof -i:9090


