

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

		fmt.Println((strings.Title(strings.ToLower(attribute))))  // strings.Title is deprecated: The rule Title uses for word boundaries does not handle Unicode punctuation properly. Use golang.org/x/text/cases instead.
		//capitalized := cases.Title(language.English).String(strings.ToLower(attribute)) // Requiere utilizar golang.org/x/text/cases (al parecer no es estandar)
		//fmt.Println("CAPITALIZED", capitalized) // Requiere utilizar golang.org/x/text/cases (al parecer no es estandar)
		fmt.Println("Capitalize alternativa nativa: ", strings.ToUpper(string(attribute[0]))+string(attribute[1:])) // toco esto para no usar mas dependencias.


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



Al parecer no es capaz de aplicar un "nested template" 


Es decir generar la data a ser llenada desde la classMetadata que forma a {{.GeneratedType}}
no es capaz de agregar la entidad

se puede corroborar que la entidad si es aplicada cuando se asigna el tag.

type {{.Entity}}Request struct {
descripcion	string	`json:"descripcion"`
precio	integer	`json:"precio"`
cantidad	integer	`json:"cantidad"`
nombre	string	`json:"nombre"`
}

type BookRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

SERA QUE SI SE PUEDE HACER ASI EN CALIENTE GENERANDO EL TAG?
PUEDO DIVIDIRLO EN DOS?

O QUIZAS DEBA TENER UN ARCHIVO TEMPORAL INTERMEDIO?



## VOY ACA


        {{.handlers_typeEntityRequest}}
tmpl es: &{fileContent 0xc00013a900 0xc000146190  }
Error al ejecutar la plantilla: template: fileContent:8:3: executing "fileContent" at <.handlers_typeEntityRequest>: handlers_typeEntityRequest is an unexported field of struct type generator.preTemplateData



## Orden de los atributos en los multiline



	go run main.go init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest | grep -A 8 -B 8 "models_UpdateErr"




## Generar Tags

Si se requiere un tag especifico

    git tag -a v0.1.1 -m "Nueva versión con el comando init"
    git push origin v0.1.1
    go run github.com/diego-all/bridge-exe@v0.1.1 init --name Diego

    git show-ref --tags
    
    git tag -d v0.1.1
    git push origin --delete v0.1.1


    git remote add origin git@github.diego-all:diego-all/evergreen-con.git  (esta funciono para evergreen)
    git remote add origin git@github.diego-all:diego-all/run-from-gh.git


    root@pho3nix:/home/diegoall/MAESTRIA_ING/CLI/PRUEBACLI# go run github.com/diego-all/run-from-gh@latest init --db postgres my-ska
    go: downloading github.com/diego-all/run-from-gh v0.0.0-20240615221752-c6170d014454
    Generando proyecto 'my-ska' con base de datos 'postgres'


    Al parecer tiene 2 comandos: my-cli-app e init

comando raiz


Usage:
  my-cli-app init [nombre del proyecto] [flags]


    go run github.com/diego-all/run-from-gh@latest init --db postgres my-ska


    go run github.com/diego-all/run-from-github@latest init --db postgres --config /ruta/al/archivo/config.json my-ska


Con este dice que le sobra un parametro

    go run github.com/diego-all/run-from-gh@latest init --db postgres /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json my-skakkkk (primera)
    
    go run github.com/diego-all/run-from-github@latest init --db postgres --config path/to/your/config.json my-ska  (segunfa)


    go run github.com/diego-all/run-from-gh@latest init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json my-skakkkk (segunda)


    go run github.com/diego-all/run-from-gh@latest init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json my-skakkkk





Con este funciona bien y dice generando proyecto

    go run main.go init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json projectTest   (MAS CERCANO)
    go run github.com/diego-all/run-from-gh@latest init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json projectTest   (MAS CERCANO)

    go run github.com/diego-all/run-from-gh@latest init --db postgres /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json

    go run github.com/diego-all/run-from-gh@latest init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json my-skakkkk


El error tendra que ver con el modulo?
con el tag?

    Ejecutar y buscar tipos

    go run main.go init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest | grep -A 10  "models_GetAllErrRowsScan"


    ACTUAL


    go run main.go init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest


    go run github.com/diego-all/run-from-gh@latest init --db postgres --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/config.json projectTest



    go run github.com/diego-all/run-from-gh@v0.1.1 init --db sqlite --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest
    go run main.go init --db sqlite --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest 