package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers.txt",

// Estructura de archivos y carpetas
var estructura = map[string]string{
	"cmd/api/handlers.go":             "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers.txt",
	"cmd/api/handlers-{{.Entity}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-{{.Entity}}.txt",
	"cmd/api/main.go":                 "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/main.txt",
	"cmd/api/routes.go":               "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/routes.txt",
	"cmd/api/util.go":                 "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/util.txt",
	"database/connection.go":          "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/connection.txt",
	"database/up.sql":                 "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/up.sql",
	"data.sqlite":                     "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/data.sqlite",
	"golang-CRUD-{{.Entity}}-API.postman_collection.json": "/base-templates",
	"go.mod":                        "/base-templates",
	"go.sum":                        "/base-templates",
	"internal/models.go":            "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/internal/models.txt",
	"internal/{{.EntityPlural}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/internal/Entities.txt",
	"product_classDiagram.png":      "/base-templates",
	"README.md":                     "/base-templates",
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
	for pathTemplate, content := range estructura {
		// Reemplazar los valores en el path
		path := strings.Replace(pathTemplate, "{{.Entity}}", class, -1)
		path = strings.Replace(path, "{{.EntityPlural}}", entityPlural, -1)
		fullPath := filepath.Join(appName, path)

		// Crear las carpetas necesarias
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			continue
		}

		// Crear el archivo
		file, err := os.Create(fullPath)
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			continue
		}
		defer file.Close()

		// Si hay contenido de plantilla, procesarlo
		if content != "" {
			tmpl, err := template.New("fileContent").Parse(content)
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
