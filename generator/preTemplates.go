package generator

import (
	"fmt"
	"strings"
)

// Los que no son Req o Res van para la DB

var TypesVars = map[string]string{
	"handlers-typeEntityRequest":     "",
	"handlers-typeEntityResponse":    "",
	"handlers-varCreateEntityModels": "",
	"handlers-varGetEntResponse":     "",
	"handlers-varUpdateEntityModels": "",
	"handlers":                       "",
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

	var auxReqRes string

	var auxCreateEntModels string //Se uso para Update, solo se adicionan 2 campos
	var auxGetEntResponse string

	var reqResTypes []string
	var createEntModels []string
	var getEntResponse []string
	var multilineAux string

	var multilineAuxCEntModels string
	var multilineAuxREntResponse string

	var handlers_typeEntityRequest string
	var handlers_typeEntityResponse string
	var handlers_varCreateEntityModels string
	var handlers_varGetEntResponse string
	var handlers_varUpdateEntityModels string

	for attribute, value := range classMetadata {

		//fmt.Printf("Clave: %s, Valor: %s\n", attribute, value)
		fmt.Println((strings.Title(strings.ToLower(attribute))))
		//capitalized := cases.Title(language.English).String(strings.ToLower(attribute)) // Requiere utilizar golang.org/x/text/cases (al parecer no es estandar)
		//fmt.Println("CAPITALIZED", capitalized) // Requiere utilizar golang.org/x/text/cases (al parecer no es estandar)
		fmt.Println("Capitalize alternativa nativa: ", strings.ToUpper(string(attribute[0]))+string(attribute[1:])) // toco esto para no usar mas dependencias.

		//auxReqRes = attribute + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
		auxReqRes = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
		auxCreateEntModels = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ":\t" + "{{.Entity}}Req." + attribute + ","
		auxGetEntResponse = attribute + ":\t" + "{{.Entity}}." + attribute + ","

		//fmt.Println("auxReqRes", auxReqRes)
		fmt.Println("AUXCREATEENTMODELS", auxCreateEntModels)
		fmt.Println("auxGetEntResponse", auxGetEntResponse)

		//Append de cada una de los atributos leidos
		reqResTypes = append(reqResTypes, auxReqRes)
		createEntModels = append(createEntModels, auxCreateEntModels)
		getEntResponse = append(getEntResponse, auxGetEntResponse)

	}
	fmt.Println("\n")
	fmt.Println("Array de reqResTypes: ", reqResTypes)
	fmt.Println("Array de createEntModels: ", createEntModels)
	fmt.Println("Array de getEntResponse: ", getEntResponse)

	fmt.Println("\n")

	// Se verticalizan
	for i, j := range reqResTypes {

		fmt.Println("Valor de i", i, "Valor de j", j)
		//multilineAux = j +"\n"
		multilineAux = multilineAux + reqResTypes[i] + "\n"
	}

	fmt.Println("multilineAux: \n ", multilineAux)
	fmt.Println("\n")

	//Tambien puede servir para el {{.Entity}}Response
	handlers_typeEntityRequest = "type {{.Entity}}Request struct {" + "\n" + multilineAux + "}"
	fmt.Println("\n")
	fmt.Println("handlers_typeEntityRequest: \n ", handlers_typeEntityRequest)

	handlers_typeEntityResponse = "type {{.Entity}}Response struct {" + "\n" + multilineAux + "}"
	fmt.Println("\n")
	fmt.Println("handlers_typeEntityResponse: \n ", handlers_typeEntityResponse)

	// Para createEntModels
	for i, j := range createEntModels {
		fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxCEntModels = multilineAuxCEntModels + createEntModels[i] + "\n"

	}

	fmt.Println("multilineAuxCEntModels: \n ", multilineAuxCEntModels)
	fmt.Println("\n")

	//para create handlers_varCreateEntityModels
	handlers_varCreateEntityModels = "var {{.Entity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varCreateEntityModels: \n ", handlers_varCreateEntityModels)

	// Para update handlers_varGetEntResponse
	handlers_varUpdateEntityModels = "var {{.Entity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "UpdatedAt:   time.Now()," + "\n" + "Id:          productID," + "\n" + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varUpdateEntityModels: \n ", handlers_varUpdateEntityModels)

	// EN construccion

	for i, j := range getEntResponse {
		fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxREntResponse = multilineAuxREntResponse + getEntResponse[i] + "\n"

	}

	fmt.Println("multilineAuxREntResponse: \n ", multilineAuxREntResponse)
	fmt.Println("\n")

	handlers_varGetEntResponse = "var {{.Entity}}Response = {{.Entity}}Response{\n" + multilineAuxREntResponse + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varGetEntResponse: \n ", handlers_varGetEntResponse)

	// Generated Types and Vars
	TypesVars["handlers-typeEntityRequest"] = handlers_typeEntityRequest
	TypesVars["handlers-typeEntityResponse"] = handlers_typeEntityResponse
	TypesVars["handlers-varCreateEntityModels"] = handlers_varCreateEntityModels
	TypesVars["handlers-varGetEntResponse"] = handlers_varGetEntResponse
	TypesVars["handlers-varUpdateEntityModels"] = handlers_varUpdateEntityModels

	fmt.Println("TIPO FINAL: ", TypesVars)

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
