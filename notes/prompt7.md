
Necesito agregar una nueva funcionalidad a este CLI pero que no sea con la funcion init sino que se llame rollback

Se pueda ejecutar de la siguiente forma:

go run main.go --rollback y haga lo siguiente:


.
├── base-templates
│   ├── cmd
│   │   └── api
│   │       ├── handlers-entity-base.txt
│   │       ├── handlers-entity-generic.txt
│   │       ├── handlers.txt
│   │       ├── main.txt
│   │       ├── routes.txt
│   │       └── util.txt
│   ├── database
│   │   ├── connection.txt
│   │   ├── data.sqlite
│   │   ├── up.sql-base.txt
│   │   └── up.sql-generic.txt
│   ├── data.sqlite
│   ├── go.mod.txt
│   ├── go.sum.txt
│   ├── internal
│   │   ├── entities-base.txt
│   │   ├── entities-generic.txt
│   │   └── models.txt
│   └── README.txt
├── books.db


Copie el contenido de cmd/api/handlers-entity-base.txt a cmd/api/handlers-entity-generic.txt ,

Tambien copue el contenido de:

database/up.sql-base.txt a database/up.sql-generic.txt


Tambien copie el contenido de:

internal/entities-base.txt a internal/entitites-generic.txt

AL final valide que los archivos equivalentes entre si sean iguales es decir los genericos sean iguales al archivo de base que es el original. e Imprima que el contenido de las pretemplates ha sido restablecido correctamente.


POdrias darme la respuesta en español y explicarme por que el archivo root.go no fue modificado

Podrias darme la respuesta en español.