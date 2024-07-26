Requiero a partir de un map[string]string (estructura) donde la key es el nombre del archivo a generar y el valor es la ruta a la base o plantilla de referencia para generar un proyecto en golang utilizando templates.

Aca esta el codigo:

package generator

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Estructura de archivos y carpetas
var estructura = map[string]string{
	"cmd/api/handlers.go":             "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers.txt",
	"cmd/api/handlers-{{.Entity}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-entity.txt",
	//"cmd/api/handlers-{{.Entity}}.go": "/base-templates/cmd/api/handlers-{{.Entity}}.txt",
	"cmd/api/main.go":        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/main.txt",
	"cmd/api/routes.go":      "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/routes.txt",
	"cmd/api/util.go":        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/util.txt",
	"database/connection.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/connection.txt",
	"database/up.sql":        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/up.sql",
	"data.sqlite":            "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/data.sqlite",
	//"golang-CRUD-{{.Entity}}-API.postman_collection.json": "/base-templates",
	"go.mod":                        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/go.mod",
	"go.sum":                        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/go.sum",
	"internal/models.go":            "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/internal/models.txt",
	"internal/{{.EntityPlural}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/internal/Entities.txt",
	//"product_classDiagram.png":      "/base-templates",
	"README.md": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/README.md",
}

// Datos para las plantillas
type TemplateData struct {
	Entity        string
	EntityPlural  string
	AppName       string
	ClassMetadata string
}

func createFolderStructure(appName string, class string) {

	// Convertir la entidad a plural
	entityPlural := class + "s"

	// TODO Class ClassMetadata
	metadata := "hola"

	// Datos para las plantillas
	data := TemplateData{
		Entity:        class,
		EntityPlural:  entityPlural,
		AppName:       appName,
		ClassMetadata: metadata,
	}

	// Crear la estructura de archivos y carpetas
	for projectFile, templatePath := range estructura {
		// Reemplazar los valores en el path
		fmt.Println("Path y Content es: ", projectFile, templatePath)

		//readTemplate(content)
		path := strings.Replace(projectFile, "{{.Entity}}", class, -1)
		path = strings.Replace(path, "{{.EntityPlural}}", entityPlural, -1)
		fullPath := filepath.Join(appName, path)

		// Crear las carpetas necesarias
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0777); err != nil {
			//if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			continue
		}

		// Crear el archivo
		file, err := os.Create(fullPath)
		fmt.Println("Creando archivo: ", file, fullPath)
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			continue
		}
		defer file.Close()

		// Si hay contenido de plantilla, procesarlo
		if templatePath != "" {

			fmt.Println("entro al if del content\n")
			tmpl, err := template.New("fileContent").Parse(templatePath)
			fmt.Println("tmpl es:", tmpl)
			if err != nil {
				fmt.Println("Error al parsear la plantilla:", err)
				continue
			}
			if err := tmpl.Execute(file, data); err != nil {
				fmt.Println("Error al ejecutar la plantilla:", err)
				continue
			}
		}
	}
	fmt.Println("Estructura de archivos y carpetas generada con éxito en la carpeta", appName)
}

func readTemplate(ruta string) {
	// leer el arreglo de bytes del archivo
	datosComoBytes, err := ioutil.ReadFile(ruta)
	fmt.Println("Leyendo template ...")
	if err != nil {
		log.Fatal(err)
	}
	// convertir el arreglo a string
	datosComoString := string(datosComoBytes)
	// imprimir el string
	fmt.Println(datosComoString)
}

Requiero modificar erste script para que me permita acceder al archivo que esta en la ruta del value y reemplazar las etiquetas que sean necesarias.

Te muestro un archivo de referencia (template) a ser modificado:
Esta en la ruta base-templates/cmd/api/routes.txt


package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/health", app.Health)

	// {{.Entity}}
	mux.Post("/EntityPlural", app.Create{{.Entity}})
	mux.Get("/EntityPlural/get/{id}", app.Get{{.Entity}})
	mux.Put("/EntityPlural/update/{id}", app.Update{{.Entity}})
	mux.Get("/EntityPlural/all", app.All{{.Entity}})
	mux.Delete("/EntityPlural/delete/{id}", app.Delete{{.Entity}})

	return mux
}

Podrias darme la respuesta en español por favor.
