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
