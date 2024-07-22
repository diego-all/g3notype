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
	"cmd/api/handlers.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers.txt",
	//"cmd/api/handlers-{{.Entity}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-entity.txt",
	"cmd/api/handlers-{{.Entity}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-entity-generic.txt",
	//"cmd/api/handlers-{{.Entity}}.go": "/home/diegoall/base-templates/cmd/api/handlers-{{.Entity}}.txt",
	//"cmd/api/handlers-{{.Entity}}.go": "/base-templates/cmd/api/handlers-{{.Entity}}.txt",
	"cmd/api/main.go":        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/main.txt",
	"cmd/api/routes.go":      "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/routes.txt",
	"cmd/api/util.go":        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/util.txt",
	"database/connection.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/connection.txt",
	"database/up.sql":        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/up.sql-generic.txt",
	//"database/up.sql":        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/up.sql",
	"data.sqlite": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/data.sqlite",
	//"golang-CRUD-{{.Entity}}-API.postman_collection.json": "/base-templates",
	"go.mod":                        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/go.mod",
	"go.sum":                        "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/go.sum",
	"internal/models.go":            "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/internal/models.txt",
	"internal/{{.EntityPlural}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/internal/entities-generic.txt", //la acabo de cambiar
	//"product_classDiagram.png":      "/base-templates",
	"README.md": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/README.md",
}

// Datos para las plantillas
// !!!!!!!!!!!!!!!!!!!!! Al parecer es necesario crear una estructura temporal ya que TemplateData no puede modificarse en tiempo de ejecucion con el fin de generar los tipos para Request y Response
// generateClassTags(class, classMetadata)
type TemplateData struct {
	Entity        string
	LowerEntity   string
	EntityPlural  string
	AppName       string
	ClassMetadata map[string]string
	GeneratedType string
}

func createFolderStructure(appName string, class string, classMetadata map[string]string) {
	//func createFolderStructure(appName string, class string, classMetadata map[string]string, generatedType string) {

	// Convertir la entidad a plural
	entityPlural := class + "s"
	fmt.Println("Imprimiendo el plural de la entidad", entityPlural)
	lowerEntity := strings.ToLower(class)

	// strings.ToUpper(string(attribute[0])) + string(attribute[1:])

	//fmt.Println("GENERATED TYPE ES desde createFolder:", generatedType)

	// TODO Class ClassMetadata
	//metadata := classMetadata

	// Datos para las plantillas
	data := TemplateData{
		Entity:       class,
		LowerEntity:  lowerEntity,
		EntityPlural: entityPlural,
		AppName:      appName,
		//ClassMetadata: classMetadata, // recien comentada
		//GeneratedType: generatedType,
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

			content, err := ioutil.ReadFile(templatePath)
			if err != nil {
				fmt.Println("Error al leer la plantilla:", err)
				continue
			}

			tmpl, err := template.New("fileContent").Parse(string(content))
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
