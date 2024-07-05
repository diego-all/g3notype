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
// !!!!!!!!!!!!!!!!!!!!! Al parecer es necesario crear una estructura temporal ya que TemplateData no puede modificarse en tiempo de ejecucion con el fin de generar los tipos para Request y Response
// generateClassTags(class, classMetadata)
type TemplateData struct {
	Entity        string
	EntityPlural  string
	AppName       string
	ClassMetadata map[string]string
	GeneratedType string
}

// quizas sea generar Tipos o algo asi, todas las estructuras que dependen de la metadata de clases (atributos)
// structs or vars// {{.Entity}}Request  {{.Entity}}Response
func generateClassTags(class string, classMetadata map[string]string) string {

	// Aca ñuego hay que jugar con mayusculas y minusculas para las entidades en tipos y variables

	fmt.Println("Desde generateClassTags")

	fmt.Println("Class metadata", classMetadata)
	longitud := len(classMetadata)
	fmt.Println("longitud del map es:", longitud)

	var aux string
	var auxCreateEntModels string
	var tagsTypes []string
	var createEntModels []string
	var multilineAux string
	var multiline string
	var multilineAuxCEntModels string
	var multilineCEntModels string
	var multilineUEntModels string
	var multilineREntResponse string

	for attribute, value := range classMetadata {

		fmt.Printf("Clave: %s, Valor: %s\n", attribute, value)
		fmt.Println((strings.Title(strings.ToLower(attribute))))
		//capitalized := cases.Title(language.English).String(strings.ToLower(attribute)) // Requiere utilizar golang.org/x/text/cases (al parecer no es estandar)
		//fmt.Println("CAPITALIZED", capitalized) // Requiere utilizar golang.org/x/text/cases (al parecer no es estandar)
		fmt.Println("alternativa nativa: ", strings.ToUpper(string(attribute[0]))+string(attribute[1:])) // toco esto para no usar mas dependencias.

		aux = attribute + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
		auxCreateEntModels = attribute + ":\t" + "{{.Entity}}Req." + attribute + ","

		fmt.Println("AUX", aux) // \t
		fmt.Println("AUXCREATEENTMODELS", auxCreateEntModels)

		tagsTypes = append(tagsTypes, aux)
		createEntModels = append(createEntModels, auxCreateEntModels)

		// strings.ToLower(s)

		// preciointeger`json:"precio"`
		// nombrestring`json:"nombre"`
		// descripcionstring`json:"descripcion"`

		// type {{.Entity}}Request struct {
		// 	Name        string  `json:"name"`
		// 	Description string  `json:"description"`
		// 	Price       float64 `json:"price"`
		// }
		// multilineAux = multilineAux + aux //Pendiente
	}
	fmt.Println("\n")
	fmt.Println("Array de tags: ", tagsTypes)
	fmt.Println("Array de createEntModels: ", createEntModels)

	// Name        string  `json:"name"`
	// Description string  `json:"description"`
	// Price       float64 `json:"price"`

	// createEntModels
	// var product = models.Product{
	// 	Name:        productReq.Name,
	// 	Description: productReq.Description,
	// 	Price:       productReq.Price,
	// }

	//UpdateEntModels
	// var product = models.Product{
	// 	Name:        productReq.Name,
	// 	Description: productReq.Description,
	// 	Price:       productReq.Price,
	// 	UpdatedAt:   time.Now(),
	// 	Id:          productID,
	// }

	//
	// var productResponse = productResponse{
	// 	Name:        product.Name,
	// 	Description: product.Description,
	// 	Price:       product.Price,
	// }

	fmt.Println("\n")

	for i, j := range tagsTypes {

		fmt.Println("Valor de i", i, "Valor de j", j)
		//multilineAux = j +"\n"
		multilineAux = multilineAux + tagsTypes[i] + "\n"
	}

	fmt.Println("multilineAux: \n ", multilineAux)
	fmt.Println("\n")

	//Tambien puede servir para el {{.Entity}}Response
	multiline = "type {{.Entity}}Request struct {" + "\n" + multilineAux + "}"
	fmt.Println("\n")
	fmt.Println("multilineFinal: \n ", multiline)

	// Para createEntModels
	for i, j := range createEntModels {
		fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxCEntModels = multilineAuxCEntModels + createEntModels[i] + "\n"

	}

	fmt.Println("multilineAuxCEntModels: \n ", multilineAuxCEntModels)
	fmt.Println("\n")

	multilineCEntModels = "var {{.Entity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "}"
	fmt.Println("\n")
	fmt.Println("var CreateEntityModels: \n ", multilineCEntModels)

	multilineUEntModels = "var {{.Entity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "UpdatedAt:   time.Now()," + "\n" + "Id:          productID," + "\n" + "}"
	fmt.Println("\n")
	fmt.Println("var UpdateEntityModels: \n ", multilineUEntModels)

	// EN construccion
	multilineREntResponse = "var {{.Entity}}Response = {{.Entity}}Response{"

	return multiline
}

func createFolderStructure(appName string, class string, classMetadata map[string]string, generatedType string) {

	// Convertir la entidad a plural
	entityPlural := class + "s"

	fmt.Println("GENERATED TYPE ES desde createFolder:", generatedType)

	// TODO Class ClassMetadata
	//metadata := classMetadata

	// Datos para las plantillas
	data := TemplateData{
		Entity:        class,
		EntityPlural:  entityPlural,
		AppName:       appName,
		ClassMetadata: classMetadata,
		GeneratedType: generatedType,
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
