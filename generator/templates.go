package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

// File and folder structure
var estructura = map[string]string{
	"cmd/api/handlers.go":             "base-templates/cmd/api/handlers.txt",
	"cmd/api/handlers-{{.Entity}}.go": "base-templates/cmd/api/handlers-entity-generic.txt",
	"cmd/api/main.go":                 "base-templates/cmd/api/main.txt",
	"cmd/api/routes.go":               "base-templates/cmd/api/routes.txt",
	"cmd/api/util.go":                 "base-templates/cmd/api/util.txt",
	"database/connection.go":          "base-templates/database/connection.txt",
	"database/up.sql":                 "base-templates/database/up.sql-generic.txt",
	"database/create-db.sh":           "base-templates/database/create-db.sh",
	//"golang-CRUD-{{.Entity}}-API.postman_collection.json": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/CRUD-API-collection.txt",
	"go.mod":      "base-templates/go.mod.txt",
	"go.sum":      "base-templates/go.sum.txt",
	"requests.md": "base-templates/requests-generic.txt",

	"internal/models.go":            "base-templates/internal/models.txt",
	"internal/{{.EntityPlural}}.go": "base-templates/internal/entities-generic.txt", //la acabo de cambiar
	"README.md":                     "base-templates/README.txt",
}

type TemplateData struct {
	Entity        string
	LowerEntity   string
	EntityPlural  string
	AppName       string
	ClassMetadata map[string]string
	GeneratedType string
}

func createFolderStructure(appName string, class string, classMetadata [][]string) {

	// Convert the entity to plural
	entityPlural := class + "s"
	//fmt.Println("Imprimiendo el plural de la entidad", entityPlural)
	lowerEntity := strings.ToLower(class)

	// Data for templates
	data := TemplateData{
		Entity:       class,
		LowerEntity:  lowerEntity,
		EntityPlural: entityPlural,
		AppName:      appName,
		//ClassMetadata: classMetadata,
		//GeneratedType: generatedType,
	}

	// Create the file and folder structure
	for projectFile, templatePath := range estructura {
		// Replace values ​​in path
		// Debugging important!
		//fmt.Println("Path y Content es: ", projectFile, templatePath)

		// Read template(content)
		path := strings.Replace(projectFile, "{{.Entity}}", class, -1)
		path = strings.Replace(path, "{{.EntityPlural}}", entityPlural, -1)
		fullPath := filepath.Join(appName, path)

		// Create the necessary folders
		dir := filepath.Dir(fullPath)
		if err := os.MkdirAll(dir, 0777); err != nil {
			//if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Println("Error al crear la carpeta:", err)
			continue
		}

		// Create the file
		file, err := os.Create(fullPath)
		//fmt.Println("Creando archivo: ", file, fullPath)
		if err != nil {
			fmt.Println("Error al crear el archivo:", err)
			continue
		}
		defer file.Close()

		// If there is template content, process it
		if templatePath != "" {

			content, err := ioutil.ReadFile(templatePath)
			if err != nil {
				fmt.Println("Error al leer la plantilla:", err)
				continue
			}

			tmpl, err := template.New("fileContent").Parse(string(content))
			// Important for debugging
			//fmt.Println("tmpl es:", tmpl)
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
	fmt.Println("Estructura de archivos y carpetas generada con éxito en la carpeta", appName, "\n")
}
