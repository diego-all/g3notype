# G3notype

Generate a API REST with Go from domain model. 


## Requirements

Golang > 1.22 and python > 3.7

### Extractor (Python)

  Needs python > 3.7

  Uso: python3 extractor/readMap.py <ruta_del_json>

  Uso: python3 extractor/readMap.py /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json


## Use

## Para su ejecución se debe especificar:


Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  init        Inicializa un nuevo proyecto
  rollback    Restaura los archivos genéricos a partir de los archivos base


init       (Iniciar la generación del scaffolding.)
--config  (Ruta al archivo de configuración JSON con el modelo)
--db sqlite  (Base de datos a utilizar)
projectTest  (Nombre del proyecto)


Flags:
  -c, --config string   Ruta del archivo JSON de configuración
  -d, --db string       Tipo de base de datos (requerido)
  -u, --dummy           Generar Dummy data usando Gemini (Requiere API Key)
  -h, --help            help for init



### Run from remote repository

    go run github.com/diego-all/run-from-gh@v0.1.1 init --db sqlite --config inputs/product.json exampleAPI


### Run local 

    go build -o g3notype .  && ./g3notype

    go run main.go init --db sqlite --dummy --config inputs/product.json projectTest

    go run github.com/diego-all/run-from-gh@latest init

    go run github.com/diego-all/run-from-gh@latest init -h



## Agregar feature Generate dummy Data

  go run main.go init --db sqlite --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest

  go run main.go init --db sqlite --config inputs/classes.json projectTest

  go run main.go init --db sqlite --dummy --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest

  go run main.go init --db sqlite --dummy --config inputs/classes.json projectTest





### Input JSON (--config)

    [
        {
          "tipo": "Product",
          "atributos": {
            "name": {
              "tipoDato": "string"
            },
            "description": {
              "tipoDato": "string"
            },
            "price": {
              "tipoDato": "int"
            },
            "quantity": {
              "tipoDato": "int"
            }
          }
        }
    ]


## Generate Dummy Data

Gemini API Key

https://aistudio.google.com/app/apikey



## Rollback (Copy generic to base)

    go run main.go rollback   (Reestablecer las Pre-Templates a su estado original)


## Generated project's structure


    ├── cmd
    │   └── api
    │       ├── handlers.go
    │       ├── handlers-Product.go
    │       ├── main.go
    │       ├── routes.go
    │       └── util.go
    ├── database
    │   ├── connection.go
    │   ├── create-db.sh
    │   └── up.sql
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── models.go
    │   └── Products.go
    ├── README.md
    └── requests.md


## Crear la base de datos

    sh create-db.sh


## Analizar el tema del repositorio

    go run ./cmd/api
    go build ./cmd/api
    go build -o productsAPI ./cmd/api

    cmd/api/handlers-Book.go:8:2: no required module provides package github.com/diego-all/books-API/internal; to add it:

    Quiza deba cambirse el nombre para poder probar y no busque modulo.

    Funciona sin repositorio. Analizar como hacer que funcione con repositorio y crearlo con git. (analizar como crea el repo git init y demas)



## Instalacion

## Comandos disponibles


## Configuracion

## Soporte

## FAQ o problemas conocidos

## Changelog

## Licencia