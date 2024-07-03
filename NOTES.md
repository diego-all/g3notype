

prompt

Requiero un programa de tipo CLI en golang utiliando la libreria spf13/cobra que permita capturar comandos para ejecutar ciertas funciones. Requiero que en el programa main exista la forma de recibir el parametro init en conjunto con el nombre del proyecto (my-ska) y el flag -db con valor ingresado por el usuario. Este programa debera permitir ser ejecutado haciendo un llamado al repositorio de github donde se encuentre la aplicacion. go run github.com/diego-all/run-from-github@latest init --db postgres my-ska Este comando init llamara la funcion generate() que estará en el paquete generator del programa. Podrias darme la respuesta en español por favor.

Podrias darme el codigo fuente por favor.


docker run --rm \
-it -p 8400-8500:8400-8500 \
-v ~/.msf4:/root/.msf4 \
-v /tmp/msf:/tmp/data \
phocean/msf



## spf13/cobra

Duda orden en los parametros.

**Argumentos Posicionales vs. Flags**

**Posicionales** son aquellos que no llevan un prefijo con guiones (--). Simplemente se pasan en el orden esperado por el comando. En tu caso, init espera un argumento posicional para el nombre del proyecto.

**Flags** son argumentos que tienen un nombre precedido por uno o dos guiones (- o --) y generalmente se usan para opciones que pueden o no ser proporcionadas. En tu comando, --db y --config son flags.




## Orden de envio de parametros de ejecucion

    go run main.go init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json projectTest

Recordar si se usa el de GitHub se debe de tener el repo actualizado. Creo que debe de estar en una rama especifica o una version especifica (tag).

    go run github.com/diego-all/run-from-gh@latest init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json projectTest


## Mejor forma de representar un sistema de archivos y carpetas



El por naturaleza seria un arbol.
Tambien se podria utilizar un JSON

Por aca en este post https://stackoverflow.com/questions/12657365/extracting-directory-hierarchy-using-go-language proponen lo siguiente:

Este man tiene algo interesante con arboles:

https://github.com/marcinwyszynski/directory_tree/blob/master/examples/find.go


Chatgpt propone un map

la estructura de datos más adecuada es un mapa (map) en Go, donde las claves representan las rutas de los archivos y los valores representan el contenido de los archivos. Esta estructura te permitirá acceder fácilmente a cada archivo y su contenido, facilitando la modificación y generación de la estructura de directorios completa.


Analisis Diego


si el contenido de los archivos es muy extenso, no es práctico almacenarlo directamente en el mapa como valores



## A generar o modificar 

por ahora sustituir

	Entity       string
	EntityPlural string
	AppName      string

En los siguientes archivos o carpetas.

    base-template/cmd/api
    handlers-{{.Entity}}.go   Al interior {{.Entity}} {{.EntityPlural}} ClassMetadata
    routes.go   {{.Entity}} {{.EntityPlural}}

    base-template/database
    up.sql  ClassMetadata  Para el final

    base-template/internal
    models.go
    {{.EntityPlural}}.go

    data.sqlite (Analizar luego)
    go.mod (Analizar luego)
    go.sum (Analizar luego) 



## Capitalizar strings

		fmt.Printf("Clave: %s, Valor: %s\n", attribute, value)
		fmt.Println((strings.Title(strings.ToLower(attribute))))
		capitalized := cases.Title(language.English).String(strings.ToLower(attribute))
		fmt.Println("CAPITALIZED", capitalized)


## Generar los tipos

Identificar todos los tipos, por ahora se hara la prueba en cmd/api/handlers-{{.Attributes}}.go

 **NO SE PUEDEN COLOCAR TAGS QUE NO HAYAN A SER SUSTITUIDOS**


type {{.Entity}}Request struct {
	//{{.Attributes}}
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type {{.Entity}}Response struct {
	//{{.Attributes}}
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}











