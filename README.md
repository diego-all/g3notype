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


## Roolback (Copy generic to base)


## Crear la base de datos


	CreateDatabase()  dummyData()

  nano create_database.sql


  sqlite3 mi_base_de_datos.db < create_database.sql

  sqlite3 mi_base_de_datos.db

  .tables
  SELECT * FROM books;



## Tiempo de ejecucion

  El tiempo de ejecución es: 13.435069ms
  El tiempo de ejecución es: 12.930049ms
  El tiempo de ejecución es: 12.957422ms
  El tiempo de ejecución es: 14.075117ms
  El tiempo de ejecución es: 12.814421ms


  dummyData.go (independiente)

  El tiempo de ejecución es: 4.799758713s
  El tiempo de ejecución es: 5.610740169s
  El tiempo de ejecución es: 3.801896691s
  El tiempo de ejecución es: 5.126581685s
  El tiempo de ejecución es: 5.268949742s
  El tiempo de ejecución es: 6.398674133s
  El tiempo de ejecución es: 5.195210352s
  El tiempo de ejecución es: 4.689731216s



  Tiempo completo de ejecucion de generator y dummyData

  El tiempo de ejecución es: 5.370131241s


## Agregar feature Generate dummy Data

  go run main.go init --db sqlite --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest

  go run main.go init --db sqlite --dummy --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest

  go run main.go init --db sqlite --dummy --config inputs/classes.json projectTest

  go run main.go rollback   (Reestablecer las Pre-Templates a su estado original)





