package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// Estructura de archivos y carpetas
var estructura = map[string]string{
	"cmd/api/handlers.go":             "",
	"cmd/api/handlers-{{.Entity}}.go": "package api\n\n// Handler para {{.Entity}}",
	"cmd/api/main.go":                 "",
	"cmd/api/routes.go":               "",
	"cmd/api/util.go":                 "",
	"database/connection.go":          "",
	"database/up.sql":                 "",
	"data.sqlite":                     "",
	"golang-CRUD-{{.Entity}}-API.postman_collection.json": "",
	"go.mod":                        "",
	"go.sum":                        "",
	"internal/models.go":            "",
	"internal/{{.EntityPlural}}.go": "package internal\n\n// Código relacionado con {{.EntityPlural}}",
	"product_classDiagram.png":      "",
	"README.md":                     "",
}

// Datos para las plantillas
type TemplateData struct {
	Entity       string
	EntityPlural string
	AppName      string
}

func createFolderStructure(appName string, class string) {
	// Solicitar el nombre de la aplicación y la entidad al usuario
	//var appName, entity string
	//fmt.Println("Ingresa el nombre de la aplicación:")
	//fmt.Scanln(&appName)
	//fmt.Println("Ingresa el nombre de la entidad (ej. product):")
	//fmt.Scanln(&entity)

	// Convertir la entidad a plural
	entityPlural := class + "s"

	// Datos para las plantillas
	data := TemplateData{
		Entity:       class,
		EntityPlural: entityPlural,
		AppName:      appName,
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
