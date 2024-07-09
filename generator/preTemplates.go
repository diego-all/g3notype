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
}

// quizas sea generar Tipos o algo asi, todas las estructuras que dependen de la metadata de clases (atributos)
// structs or vars// {{.Entity}}Request  {{.Entity}}Response
// INTENTAR GENERANDO LOS TIPOS PRIMERO DE HANDLERS, LUEGO PARA MODELS U OTROS DE SER NECESARIO MANEJARLOS EN UN MAP

// func generateClassTags(class string, classMetadata map[string]string) (string) {
func generateClassTags(class string, classMetadata map[string]string) map[string]string {

	fmt.Println("Desde generateClassTags")

	fmt.Println("Class metadata", classMetadata)
	longitud := len(classMetadata)
	fmt.Println("longitud del map es:", longitud)
	fmt.Println("\n")

	var auxReqRes string
	var auxCreateEntModels string //Se uso para Update, solo se adicionan 2 campos
	var auxGetEntResponse string

	var reqResTypes []string
	var createEntModels []string
	var getEntResponse []string

	var multilineAuxReqResTypes string
	var multilineAuxCEntModels string
	var multilineAuxGEntResponse string

	var handlers_typeEntityRequest string
	var handlers_typeEntityResponse string
	var handlers_varCreateEntityModels string
	var handlers_varGetEntResponse string
	var handlers_varUpdateEntityModels string

	for attribute, value := range classMetadata {

		//fmt.Printf("Clave: %s, Valor: %s\n", attribute, value)
		fmt.Println("Capitalize alternativa nativa: ", strings.ToUpper(string(attribute[0]))+string(attribute[1:])) // toco esto para no usar mas dependencias.

		//auxReqRes = attribute + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
		auxReqRes = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
		auxCreateEntModels = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ":\t" + "{{.entity}}Req." + strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ","
		auxGetEntResponse = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ":\t" + "{{.entity}}." + strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ","

		//fmt.Println("auxReqRes", auxReqRes)
		//fmt.Println("auxCreateEntModels", auxCreateEntModels)
		//fmt.Println("AuxGetEntResponse", auxGetEntResponse)

		//Append de cada una de los atributos leidos
		reqResTypes = append(reqResTypes, auxReqRes)
		createEntModels = append(createEntModels, auxCreateEntModels)
		getEntResponse = append(getEntResponse, auxGetEntResponse)
	}

	// fmt.Println("Array de reqResTypes: ", reqResTypes)
	// fmt.Println("Array de createEntModels: ", createEntModels)
	// fmt.Println("Array de getEntResponse: ", getEntResponse)
	fmt.Println("\n")

	// Se verticalizan , creo que quedarian mejor con un while
	for i, _ := range reqResTypes {
		//fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxReqResTypes = multilineAuxReqResTypes + reqResTypes[i] + "\n"
	}

	//fmt.Println("multilineAuxReqResTypes: \n ", multilineAuxReqResTypes+"\n")
	//fmt.Println("\n")

	handlers_typeEntityRequest = "type {{.Entity}}Request struct {" + "\n" + multilineAuxReqResTypes + "}"
	//fmt.Println("\n")
	fmt.Println("handlers_typeEntityRequest: \n ", handlers_typeEntityRequest)
	fmt.Println("\n")

	handlers_typeEntityResponse = "type {{.Entity}}Response struct {" + "\n" + multilineAuxReqResTypes + "}"
	fmt.Println("handlers_typeEntityResponse: \n ", handlers_typeEntityResponse)
	fmt.Println("\n")

	// Para createEntModels
	for i, _ := range createEntModels {
		//fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxCEntModels = multilineAuxCEntModels + createEntModels[i] + "\n"
	}

	//fmt.Println("multilineAuxCEntModels: \n ", multilineAuxCEntModels)
	//fmt.Println("\n")

	//para create handlers_varCreateEntityModels
	handlers_varCreateEntityModels = "var {{.entity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varCreateEntityModels: \n ", handlers_varCreateEntityModels)

	// Para update handlers_varUpdateEntResponse
	handlers_varUpdateEntityModels = "var {{.entity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "UpdatedAt:   time.Now()," + "\n" + "Id:          productID," + "\n" + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varUpdateEntityModels: \n ", handlers_varUpdateEntityModels)

	for i, _ := range getEntResponse {
		//fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxGEntResponse = multilineAuxGEntResponse + getEntResponse[i] + "\n"
	}

	//fmt.Println("multilineAuxGEntResponse: \n ", multilineAuxGEntResponse)
	fmt.Println("\n")

	handlers_varGetEntResponse = "var {{.entity}}Response = {{.entity}}Response{\n" + multilineAuxGEntResponse + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varGetEntResponse: \n ", handlers_varGetEntResponse)

	fmt.Println("\n")
	// Generated Types and Vars
	TypesVars["handlers-typeEntityRequest"] = handlers_typeEntityRequest
	TypesVars["handlers-typeEntityResponse"] = handlers_typeEntityResponse
	TypesVars["handlers-varCreateEntityModels"] = handlers_varCreateEntityModels
	TypesVars["handlers-varGetEntResponse"] = handlers_varGetEntResponse
	TypesVars["handlers-varUpdateEntityModels"] = handlers_varUpdateEntityModels

	fmt.Println("TIPO FINAL: ", TypesVars)
	fmt.Println("\n")

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
