package generator

import (
	"fmt"
	"strings"
)

// Los que no son Req o Res van para la DB

var TypesVars = map[string]string{
	"handlers-typeEntityRequest":  "",
	"handlers-typeEntityResponse": "",
	"handlers-createEntityModels": "",
	"handlers-readEntResponse":    "",
	"handlers-updateEntityModels": "",
	"handlers":                    "",
}

// quizas sea generar Tipos o algo asi, todas las estructuras que dependen de la metadata de clases (atributos)
// structs or vars// {{.Entity}}Request  {{.Entity}}Response
// INTENTAR GENERANDO LOS TIPOS PRIMERO DE HANDLERS, LUEGO PARA MODELS U OTROS DE SER NECESARIO MANEJARLOS EN UN MAP

// func generateClassTags(class string, classMetadata map[string]string) (string) {
func generateClassTags(class string, classMetadata map[string]string) map[string]string {

	// Aca luego hay que jugar con mayusculas y minusculas para las entidades en tipos y variables

	fmt.Println("Desde generateClassTags")

	fmt.Println("Class metadata", classMetadata)
	longitud := len(classMetadata)
	fmt.Println("longitud del map es:", longitud)

	var aux string
	var auxCreateEntModels string //Se uso para Update, solo se adicionan 2 campos
	var auxReadEntResponse string
	var tagsTypes []string
	var createEntModels []string
	var readEntResponse []string
	var multilineAux string
	var multiline string
	var multilineAuxCEntModels string
	var multilineAuxREntResponse string
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
		auxReadEntResponse = attribute + ":\t" + "{{.Entity}}." + attribute + ","

		fmt.Println("AUX", aux) // \t
		fmt.Println("AUXCREATEENTMODELS", auxCreateEntModels)
		fmt.Println("AUXREADENTRESPONSE", auxReadEntResponse)

		tagsTypes = append(tagsTypes, aux)
		createEntModels = append(createEntModels, auxCreateEntModels)
		readEntResponse = append(readEntResponse, auxReadEntResponse)

	}
	fmt.Println("\n")
	fmt.Println("Array de tags: ", tagsTypes)
	fmt.Println("Array de createEntModels: ", createEntModels)
	fmt.Println("Array de readEntResponse: ", readEntResponse)

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

	//para create
	multilineCEntModels = "var {{.Entity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "}"
	fmt.Println("\n")
	fmt.Println("var CreateEntityModels: \n ", multilineCEntModels)

	// Para update
	multilineUEntModels = "var {{.Entity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "UpdatedAt:   time.Now()," + "\n" + "Id:          productID," + "\n" + "}"
	fmt.Println("\n")
	fmt.Println("var UpdateEntityModels: \n ", multilineUEntModels)

	// EN construccion

	for i, j := range readEntResponse {
		fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxREntResponse = multilineAuxREntResponse + readEntResponse[i] + "\n"

	}

	fmt.Println("multilineAuxREntResponse: \n ", multilineAuxREntResponse)
	fmt.Println("\n")

	multilineREntResponse = "var {{.Entity}}Response = {{.Entity}}Response{\n" + multilineAuxREntResponse + "}"
	fmt.Println("\n")
	fmt.Println("var ReadEntityResponse: \n ", multilineREntResponse)

	return TypesVars
	//return multiline // antes retornaba el primer type EntityRequest
}

// func modifyBaseTemplates(class string, classMetadata map[string]string) {

// 	fmt.Println("Desde modifyBaseTemplates")

// 	fmt.Println("Class metadata", classMetadata)
// 	longitud := len(classMetadata)
// 	fmt.Println("longitud del map es:", longitud)

// 	// var aux string
// 	// var tagsTypes []string
// 	// var multilineAux string
// 	// var multiline string

// 	for attribute, value := range classMetadata {

// 		fmt.Printf("Clave: %s, Valor: %s\n", attribute, value)

// 	}

// 	preData := TemplateData{
// 		Entity: class,
// 		//GeneratedType: generatedType,
// 	}

// 	fmt.Println(preData)

// }

// preData := TemplateData{
// 	Entity:        class,
// 	EntityPlural:  entityPlural,
// 	AppName:       appName,
// 	ClassMetadata: classMetadata,
// 	GeneratedType: generatedType,
// }

// type {{.Entity}}Request struct {
// 	nombre	string	`json:"nombre"`
// 	descripcion	string	`json:"descripcion"`
// 	precio	integer	`json:"precio"`
// 	cantidad	integer	`json:"cantidad"`
// 	}
