package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

// Los que no son Req o Res van para la DB

var TypesVars = map[string]string{
	"handlers-typeEntityRequest":     "",
	"handlers-typeEntityResponse":    "",
	"handlers-varCreateEntityModels": "",
	"handlers-varGetEntResponse":     "",
	"handlers-varUpdateEntityModels": "",
	"handlers-payloadCreateResponse": "",
	"handlers-payloadUpdateResponse": "",

	//Database_DDL_statement
	"database-DDL-statement": "",

	// EntityModels
	"models-typeEntityStruct":  "",
	"models-InsertStmt":        "",
	"models-InsertErr":         "",
	"models-GetOneQuery":       "",
	"models-GetOneErr":         "",
	"models-UpdateStmt":        "",
	"models-UpdateErr":         "",
	"models-GetAllQuery":       "",
	"models-GetAllErrRowsScan": "",
	"models-DeleteStmt":        "", // validar si realmente es necesario
}

// RECORDAR EL CAMBIASO DE BASE A GENERIC SOLO FUR POR ORDEN ALFABETICO PERO AL KEY NO SE CAMBIO, DEBERIA SER entity-generic.go

var preTemplates = map[string]string{
	"cmd/api/handlers-entity-base.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-entity-generic.txt",

	"database/up.sql": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/up.sql-generic.txt",

	"internal/entities-base.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/internal/entities-generic.txt",

	//"cmd/api/handlers-entity-base.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-entity-base.txt",
	//"cmd/api/handlers-{{.Entity}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-entity.txt",
}

type PreTemplateData struct {
	Handlers_typeEntityRequest     string
	Handlers_typeEntityResponse    string
	Handlers_varCreateEntityModels string
	Handlers_varGetEntResponse     string
	Handlers_varUpdateEntityModels string
	Handlers_payloadCreateResponse string
	Handlers_payloadUpdateResponse string

	//TypesVars["handlers-varUpdateEntityModels"] = Database_DDL_statement
	//"database-DDL-statement": "",
	Database_DDL_statement string

	Models_typeEntityStruct  string
	Models_InsertStmt        string
	Models_InsertErr         string
	Models_GetOneQuery       string
	Models_GetOneErr         string
	Models_UpdateStmt        string
	Models_UpdateErr         string
	Models_GetAllQuery       string
	Models_GetAllErrRowsScan string
	Models_DeleteStmt        string

	Entity      string
	LowerEntity string
}

// quizas sea generar Tipos o algo asi, todas las estructuras que dependen de la metadata de clases (atributos)
// structs or vars// {{.Entity}}Request  {{.Entity}}Response
// INTENTAR GENERANDO LOS TIPOS PRIMERO DE HANDLERS, LUEGO PARA MODELS U OTROS DE SER NECESARIO MANEJARLOS EN UN MAP

// func generateClassTags(class string, classMetadata map[string]string) (string) {
// func generateClassTags(class string, classMetadata map[string]string) map[string]string {
func generateClassTags(class string, classMetadata [][]string) map[string]string {

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
	var handlers_payloadCreateResponse string
	var handlers_payloadUpdateResponse string

	var naturalId string

	// REPETIDO
	// for attribute, value := range classMetadata {

	// 	//fmt.Printf("Clave: %s, Valor: %s\n", attribute, value)
	// 	//fmt.Println("Capitalize alternativa nativa: ", strings.ToUpper(string(attribute[0]))+string(attribute[1:])) // toco esto para no usar mas dependencias.

	// 	//auxReqRes = attribute + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
	// 	auxReqRes = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
	// 	auxCreateEntModels = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ":\t" + "{{.LowerEntity}}Req." + strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ","
	// 	auxGetEntResponse = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ":\t" + "{{.LowerEntity}}." + strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ","

	// 	//fmt.Println("auxReqRes", auxReqRes)
	// 	//fmt.Println("auxCreateEntModels", auxCreateEntModels)
	// 	//fmt.Println("AuxGetEntResponse", auxGetEntResponse)

	// 	//Append de cada una de los atributos leidos
	// 	reqResTypes = append(reqResTypes, auxReqRes)
	// 	createEntModels = append(createEntModels, auxCreateEntModels)
	// 	getEntResponse = append(getEntResponse, auxGetEntResponse)
	// }

	for k, attribute := range classMetadata {

		attributeName := attribute[0]
		attributeType := attribute[1]

		//fmt.Printf("Clave: %s, Valor: %s\n", attribute, value)
		//fmt.Println("Capitalize alternativa nativa: ", strings.ToUpper(string(attribute[0]))+string(attribute[1:])) // toco esto para no usar mas dependencias.

		//auxReqRes = attribute + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
		// auxReqRes = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + "\t" + value + "\t" + "`json:\"" + attribute + "\"`"
		// auxCreateEntModels = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ":\t" + "{{.LowerEntity}}Req." + strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ","
		// auxGetEntResponse = strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ":\t" + "{{.LowerEntity}}." + strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + ","

		auxReqRes = "\t" + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + "\t" + attributeType + "\t" + "`json:\"" + attributeName + "\"`"
		auxCreateEntModels = "\t" + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ":\t" + "{{.LowerEntity}}Req." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","
		auxGetEntResponse = "\t" + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ":\t" + "{{.LowerEntity}}." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","

		//fmt.Println("auxReqRes", auxReqRes)
		//fmt.Println("auxCreateEntModels", auxCreateEntModels)
		//fmt.Println("AuxGetEntResponse", auxGetEntResponse)

		//Append de cada una de los atributos leidos
		reqResTypes = append(reqResTypes, auxReqRes)
		createEntModels = append(createEntModels, auxCreateEntModels)
		getEntResponse = append(getEntResponse, auxGetEntResponse)

		if k == 0 {
			naturalId = string(strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]))
		}
	}

	fmt.Println(naturalId)

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

	handlers_typeEntityRequest = "type {{.LowerEntity}}Request struct {" + "\n" + multilineAuxReqResTypes + "}"
	//fmt.Println("\n")
	fmt.Println("handlers_typeEntityRequest: \n ", handlers_typeEntityRequest)
	fmt.Println("\n")

	handlers_typeEntityResponse = "type {{.LowerEntity}}Response struct {" + "\n" + multilineAuxReqResTypes + "}"
	fmt.Println("handlers_typeEntityResponse: \n ", handlers_typeEntityResponse)
	fmt.Println("\n")

	// Para createEntModels
	for i, _ := range createEntModels {
		//fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxCEntModels = multilineAuxCEntModels + createEntModels[i] + "\n"
	}

	//fmt.Println("multilineAuxCEntModels: \n ", multilineAuxCEntModels)
	//fmt.Println("\n")
	// MINUSCULA
	//para create handlers_varCreateEntityModels
	handlers_varCreateEntityModels = "var {{.LowerEntity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varCreateEntityModels: \n ", handlers_varCreateEntityModels)

	// Para update handlers_varUpdateEntResponse
	handlers_varUpdateEntityModels = "var {{.LowerEntity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "\t" + "UpdatedAt:   time.Now()," + "\n \t" + "Id:          {{.LowerEntity}}ID," + "\n" + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varUpdateEntityModels: \n ", handlers_varUpdateEntityModels)

	for i, _ := range getEntResponse {
		//fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxGEntResponse = multilineAuxGEntResponse + getEntResponse[i] + "\n"
	}

	//fmt.Println("multilineAuxGEntResponse: \n ", multilineAuxGEntResponse)
	fmt.Println("\n")

	handlers_varGetEntResponse = "var {{.LowerEntity}}Response = {{.LowerEntity}}Response{\n" + multilineAuxGEntResponse + "}"
	fmt.Println("\n")
	fmt.Println("handlers_varGetEntResponse: \n ", handlers_varGetEntResponse)

	fmt.Println("\n")

	handlers_payloadCreateResponse = "payload = jsonResponse{\n" + "\t    Error:   false,\n" + "\t    Message: \"{{.Entity}} successfully created\",\n" + "\t    Data:    envelope{\"book\": {{.LowerEntity}}." + naturalId + "},\n" + "\t}"

	fmt.Println("handlers_payloadCreateResponse: \n ", handlers_payloadCreateResponse)

	fmt.Println("\n")

	handlers_payloadUpdateResponse = "payload = jsonResponse{\n" + "\t    Error:   false,\n" + "\t    Message: \"{{.Entity}} successfully updated\",\n" + "\t    Data:    envelope{\"book\": {{.LowerEntity}}." + naturalId + "},\n" + "\t}"

	fmt.Println("handlers_payloadUpdateResponse: \n ", handlers_payloadUpdateResponse)

	fmt.Println("\n")

	// Generated Types and Vars
	TypesVars["handlers-typeEntityRequest"] = handlers_typeEntityRequest
	TypesVars["handlers-typeEntityResponse"] = handlers_typeEntityResponse
	TypesVars["handlers-varCreateEntityModels"] = handlers_varCreateEntityModels
	TypesVars["handlers-varGetEntResponse"] = handlers_varGetEntResponse
	TypesVars["handlers-varUpdateEntityModels"] = handlers_varUpdateEntityModels
	TypesVars["handlers-payloadCreateResponse"] = handlers_payloadCreateResponse
	TypesVars["handlers-payloadUpdateResponse"] = handlers_payloadUpdateResponse

	//fmt.Println("TIPO FINAL: ", TypesVars)
	//fmt.Println("\n")

	return TypesVars
	//return multiline // antes retornaba el primer type EntityRequest
}

func modifyBaseTemplates(preGeneratedTypes map[string]string) {

	fmt.Println("\n")
	fmt.Println("PREGENERATED TYPES en modifyBaseTemplates: \n", preGeneratedTypes)
	fmt.Println("LONGITUD DE PREREGENERATED TYPES en modifyBaseTemplates", len(preGeneratedTypes))

	count := 0
	for _, value := range preGeneratedTypes {
		// Verifica si el valor no es vacío. Ajusta la condición según el tipo de valor esperado.
		if value != "" { // Esto es solo un ejemplo para valores string. Ajusta según tu tipo de dato.
			count++
		}
	}
	fmt.Println("Número de keys llenas en preGeneratedTypes:", count)

	fmt.Println("\n")

	// //Error al ejecutar la plantilla: template: fileContent:8:2: executing "fileContent" at <.handlers_typeEntityRequest>: handlers_typeEntityRequest is an unexported field of struct type generator.preTemplateData
	preData := PreTemplateData{
		Handlers_typeEntityRequest:     preGeneratedTypes["handlers-typeEntityRequest"],
		Handlers_typeEntityResponse:    preGeneratedTypes["handlers-typeEntityResponse"],
		Handlers_varCreateEntityModels: preGeneratedTypes["handlers-varCreateEntityModels"],
		Handlers_varGetEntResponse:     preGeneratedTypes["handlers-varGetEntResponse"],
		Handlers_varUpdateEntityModels: preGeneratedTypes["handlers-varUpdateEntityModels"],
		Handlers_payloadCreateResponse: preGeneratedTypes["handlers-payloadCreateResponse"],
		Handlers_payloadUpdateResponse: preGeneratedTypes["handlers-payloadUpdateResponse"],

		Database_DDL_statement: preGeneratedTypes["database-DDL-statement"],

		Models_typeEntityStruct:  preGeneratedTypes["models-typeEntityStruct"],
		Models_InsertStmt:        preGeneratedTypes["models-InsertStmt"],
		Models_InsertErr:         preGeneratedTypes["models-InsertErr"],
		Models_GetOneQuery:       preGeneratedTypes["models-GetOneQuery"],
		Models_GetOneErr:         preGeneratedTypes["models-GetOneErr"],
		Models_UpdateStmt:        preGeneratedTypes["models-UpdateStmt"],
		Models_UpdateErr:         preGeneratedTypes["models-UpdateErr"],
		Models_GetAllQuery:       preGeneratedTypes["models-GetAllQuery"],
		Models_GetAllErrRowsScan: preGeneratedTypes["models-GetAllErrRowsScan"],
		Models_DeleteStmt:        preGeneratedTypes["models-DeleteStmt"], // validar si realmente es necesario

		// quizas pueda ser {{.UpperEntity}}
		Entity: "{{.Entity}}",
		//entity:  "{{.entity}}",   // NO funciona con minusculas seguir indagando
		LowerEntity: "{{.LowerEntity}}",
	}

	//fmt.Println(preData)

	// 	for attribute, value := range preGeneratedTypes {
	for projectFile, templatePath := range preTemplates {

		//fmt.Printf("Clave: %s, Valor: %s\n", projectFile, templatePath)
		fmt.Println("Path y Content es: ", projectFile, templatePath)

		// Si hay contenido de plantilla, procesarlo
		if templatePath != "" {

			content, err := ioutil.ReadFile(templatePath)
			if err != nil {
				fmt.Println("Error al leer la plantilla:", err)
				continue
			}

			//fmt.Println("CONTENT:", string(content))

			tmpl, err := template.New("fileContent").Parse(string(content))
			fmt.Println("tmpl es:", tmpl)
			if err != nil {
				fmt.Println("Error al parsear la plantilla:", err)
				continue
			}

			// Crear o abrir el archivo de destino en modo escritura
			file, err := os.Create(templatePath)
			//file, err := os.Create(projectFile)
			if err != nil {
				fmt.Println("Error al crear el archivo:", err)
				continue
			}
			defer file.Close()

			if err := tmpl.Execute(file, preData); err != nil {
				fmt.Println("Error al ejecutar la plantilla:", err)
				continue
			}
			fmt.Println("Archivo procesado correctamente:", projectFile)

		}

	}

	fmt.Println("\n")

}

// Generate Create table
func generateDDLStatement(class string, classMetadata [][]string) string {

	fmt.Println("Desde generateDDLStatement", class)

	fmt.Println("Class metadata", classMetadata)
	longitud := len(classMetadata)
	fmt.Println("longitud del map es:", longitud)
	fmt.Println("\n")

	var auxDDL string
	var ddlStatement []string
	var sqliteValue string
	var multilineAuxDDLStatement string
	var Database_DDL_statement string

	// for _, attribute := range classMetadata {

	// 	attributeName := attribute[0]
	// 	attributeType := attribute[1]

	// }

	for _, data := range classMetadata {
		//fmt.Printf("Clave: %s, Valor: %s\n", iterator, data)

		attributeName := data[0]
		attributeType := data[1]

		//fmt.Println(attributeName, attributeType)

		//fmt.Println("Capitalize alternativa nativa: ", strings.ToUpper(string(attribute[0]))+string(attribute[1:])) // toco esto para no usar mas dependencias.

		switch attributeType {
		case "integer":
			fmt.Println("INTEGER")
			sqliteValue = "INTEGER"
		case "string":
			//fmt.Println("VARCHAR")
			sqliteValue = "VARCHAR(100)"
		case "":
			fmt.Println("OTRO CASO")

		}

		auxDDL = "\t" + attributeName + " " + sqliteValue + " " + "NOT NULL,"
		ddlStatement = append(ddlStatement, auxDDL)
	}

	//fmt.Println("Array de ddlStatement: ", ddlStatement)
	fmt.Println("\n")

	fmt.Println("\n")

	// Se verticalizan , creo que quedarian mejor con un while
	for i, _ := range ddlStatement {
		//fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxDDLStatement = multilineAuxDDLStatement + ddlStatement[i] + "\n"
	}

	//fmt.Println("multilineAuxDDLStatement: ", multilineAuxDDLStatement)

	Database_DDL_statement = "CREATE TABLE IF NOT EXISTS {{.LowerEntity}}s (\n \t" + "id INTEGER PRIMARY KEY AUTOINCREMENT,\n" + multilineAuxDDLStatement + "\t" + "created_at TIMESTAMP DEFAULT DATETIME,\n \t" + "updated_at TIMESTAMP NOT NULL\n \t" + ");"

	TypesVars["database-DDL-statement"] = Database_DDL_statement

	//fmt.Println("database_DDL_statement ES:", database_DDL_statement)
	return Database_DDL_statement
}

// Generate Create table
// func generateDDLStatement(class string, classMetadata [][]string) string {

// 	fmt.Println("Desde generateDDLStatement", class)

// 	fmt.Println("Class metadata", classMetadata)
// 	longitud := len(classMetadata)
// 	fmt.Println("longitud del map es:", longitud)
// 	fmt.Println("\n")

// 	var auxDDL string
// 	var ddlStatement []string
// 	var sqliteValue string
// 	var multilineAuxDDLStatement string
// 	var Database_DDL_statement string

// 	for attribute, value := range classMetadata {
// 		//fmt.Printf("Clave: %s, Valor: %s\n", attribute, value)

// 		//fmt.Println("Capitalize alternativa nativa: ", strings.ToUpper(string(attribute[0]))+string(attribute[1:])) // toco esto para no usar mas dependencias.

// 		// if value == "integer" {
// 		// 	fmt.Println("EL VALOR ES INTEGER")
// 		// }

// 		switch value {
// 		case "integer":
// 			//fmt.Println("INTEGER")
// 			sqliteValue = "INTEGER"
// 		case "string":
// 			//fmt.Println("VARCHAR")
// 			sqliteValue = "VARCHAR(100)"
// 		case "":
// 			fmt.Println("OTRO CASO")

// 		}

// 		auxDDL = attribute + " " + sqliteValue + " " + "NOT NULL,"
// 		ddlStatement = append(ddlStatement, auxDDL)
// 	}

// 	//fmt.Println("Array de ddlStatement: ", ddlStatement)
// 	fmt.Println("\n")

// 	fmt.Println("\n")

// 	// Se verticalizan , creo que quedarian mejor con un while
// 	for i, _ := range ddlStatement {
// 		//fmt.Println("Valor de i", i, "Valor de j", j)
// 		multilineAuxDDLStatement = multilineAuxDDLStatement + "\t" + ddlStatement[i] + "\n"
// 	}

// 	//fmt.Println("multilineAuxDDLStatement: ", multilineAuxDDLStatement)

// 	Database_DDL_statement = "CREATE TABLE IF NOT EXISTS {{.LowerEntity}}s (\n" + "\t" + "id INTEGER PRIMARY KEY AUTOINCREMENT,\n" + multilineAuxDDLStatement + "\t" + "created_at TIMESTAMP DEFAULT DATETIME,\n \t" + "updated_at TIMESTAMP NOT NULL\n \t" + ");"

// 	TypesVars["database-DDL-statement"] = Database_DDL_statement

// 	//fmt.Println("database_DDL_statement ES:", database_DDL_statement)
// 	return Database_DDL_statement
// }

func generateEntityModels(class string, classMetadata [][]string) map[string]string {
	fmt.Println("Desde generateEntityModels", class)

	fmt.Println("Class metadata", classMetadata)
	longitud := len(classMetadata)
	fmt.Println("longitud del map es:", longitud)
	fmt.Println("\n")

	// Generate models-typeEntityStruct
	// Pilas con "name,omitempty"`
	var auxTypeEntityStruct, multilineAuxTypeEntityStructs, models_typeEntityStruct string
	var auxInsertStmt, models_InsertStmt string
	var auxInsertErr, multilineAuxInsertErr, models_InsertErr string
	var auxGetOneQuery, models_GetOneQuery string // similar to auxGetOneErr validar
	var auxGetOneErr, multilineAuxGetOneErr, models_GetOneErr string
	var auxUpdateStmt, multilineAuxUpdateStmt, models_UpdateStmt string
	var auxUpdateErr, multilineAuxUpdateErr, models_UpdateErr string
	var models_GetAllQuery string // similar to auxGetOneErr validar
	var models_GetAllErrRowsScan string

	var typeEntityStructs []string
	var InsertErrs []string
	var GetOneErrs []string
	var UpdateStmt []string
	var UpdateErr []string

	i := 1
	for _, data := range classMetadata {

		attributeName := data[0]
		attributeType := data[1]

		//fmt.Printf("Iterador i: %d\n", i)
		auxTypeEntityStruct = strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + "\t" + attributeType + "\t" + "`json:\"" + attributeName + "\"`"
		auxInsertStmt = auxInsertStmt + attributeName + ", "
		//auxInsertErr = auxInsertErr + "{{.LowerEntity}}." + strings.ToUpper(string(attribute[0])) + string(attribute[1:]) + "\t"
		auxInsertErr = "\t" + "{{.LowerEntity}}." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","
		auxGetOneQuery = auxGetOneQuery + attributeName + ", "
		auxGetOneErr = "\t" + "&{{.LowerEntity}}." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","
		auxUpdateStmt = "\t" + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + " = $" + strconv.Itoa(i) + ","
		auxUpdateErr = "\t" + "{{.LowerEntity}}." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","

		//fmt.Println("auxTypeEntityStruct: ", auxTypeEntityStruct)
		typeEntityStructs = append(typeEntityStructs, auxTypeEntityStruct)
		InsertErrs = append(InsertErrs, auxInsertErr)
		GetOneErrs = append(GetOneErrs, auxGetOneErr)
		UpdateStmt = append(UpdateStmt, auxUpdateStmt)
		UpdateErr = append(UpdateErr, auxUpdateErr)
		i++
	}
	//fmt.Println("Array de typeEntityStructs: ", typeEntityStructs)
	//fmt.Println("\n")
	// fmt.Println("TRINauxInsertErr: ", auxInsertErr)
	// fmt.Println("auxUpdateStmt ", auxUpdateStmt)
	// fmt.Println("auxUpdateErr", auxUpdateErr)

	// Se verticalizan , creo que quedarian mejor con un while
	for i, _ := range typeEntityStructs {
		//fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxTypeEntityStructs = multilineAuxTypeEntityStructs + "\t" + typeEntityStructs[i] + "\n"
	}

	//fmt.Println("multilineAuxTypeEntityStructs: \n ", multilineAuxTypeEntityStructs+"\n")
	fmt.Println("\n")

	models_typeEntityStruct = "type {{.Entity}} struct {" + "\n \t" + "Id	int	`json:\"id\"`" + "\n" + multilineAuxTypeEntityStructs + "\t" + "CreatedAt   time.Time `json:\"created_at\"`" + "\n \t" + "UpdatedAt   time.Time `json:\"updated_at\"`" + "\n" + "}"
	//fmt.Println("\n")
	fmt.Println("models_typeEntityStruct: \n ", models_typeEntityStruct)
	fmt.Println("\n")

	// Generate models-InsertStmt
	generateStmtValues(longitud + 2) // created_at, updated_at
	models_InsertStmt = "\tstmt := `insert into {{.LowerEntity}}s (" + auxInsertStmt + "created_at, updated_at)\n \t" + "values (" + generateStmtValues(longitud+2) + ")" + " returning  id`"
	fmt.Println("models_InsertStmt es: \n", models_InsertStmt)
	fmt.Println("\n")

	// Generate models-InsertErr
	for i, _ := range InsertErrs {
		//fmt.Println("Valor de i", i, "Valor de j", j)

		multilineAuxInsertErr = multilineAuxInsertErr + InsertErrs[i] + "\n"
	}

	//fmt.Println("multilineAuxInsertErr: \n ", multilineAuxInsertErr+"\n")
	fmt.Println("\n")

	models_InsertErr = "err := db.QueryRowContext(ctx, stmt," + "\n" + multilineAuxInsertErr + "\t" + "time.Now()," + "\n" + "\t" + "time.Now()," + "\n" + ").Scan(&newID)"

	//var auxInsertErr, InsertErr string
	//fmt.Println(auxInsertErr, InsertErr)

	fmt.Println("models_InsertErr es: \n ", models_InsertErr)
	fmt.Println("\n")

	// Generate models-GetOneQuery

	models_GetOneQuery = "\tquery := `select id, " + auxGetOneQuery + "created_at, updated_at from {{.LowerEntity}}s where id = $1`"
	// query := `select id, name, description, price, created_at, updated_at from products where id = $1`
	// query := `select id, nombre, descripcion, precio, cantidad, random, created_at, updated_at from {{.LowerEntity}} where id = $1`
	fmt.Println("models_GetOneQuery es: \n ", models_GetOneQuery)

	// Generate models-GetOneErr
	for i, _ := range GetOneErrs {
		///fmt.Println("Valor de i", i, "Valor de j", j)

		multilineAuxGetOneErr = multilineAuxGetOneErr + GetOneErrs[i] + "\n"
	}

	//fmt.Println("multilineAuxGetOneErr: \n ", multilineAuxGetOneErr+"\n")
	fmt.Println("\n")

	models_GetOneErr = "err := row.Scan(" + "\n" + multilineAuxGetOneErr + "\t" + "&{{.LowerEntity}}.Id," + "\n \t" + "&{{.LowerEntity}}.CreatedAt," + "\n" + "\t" + "&{{.LowerEntity}}.UpdatedAt," + "\n" + ")"
	fmt.Println("models_GetOneErr es: \n ", models_GetOneErr)

	fmt.Println("\n")

	// Generate models-GetAllErrRowsScan
	models_GetAllErrRowsScan = "err := rows.Scan(" + "\n" + multilineAuxGetOneErr + "\t" + "&{{.LowerEntity}}.Id," + "\n \t" + "&{{.LowerEntity}}.CreatedAt," + "\n" + "\t" + "&{{.LowerEntity}}.UpdatedAt," + "\n" + ")"
	fmt.Println("models_GetAllErrRowsScan es: \n ", models_GetAllErrRowsScan)

	fmt.Println("\n")

	// Generate models-UpdateStmt
	for i, _ := range UpdateStmt {
		//fmt.Println("Valor de i", i, "Valor de j", j)

		multilineAuxUpdateStmt = multilineAuxUpdateStmt + UpdateStmt[i] + "\n"
	}

	models_UpdateStmt = "stmt := `update {{.LowerEntity}} set" + "\n" + " " + multilineAuxUpdateStmt + "\t" + "updated_at = $" + strconv.Itoa(i+1) + "\n \t" + "where id = $" + strconv.Itoa(i+2) + "`"

	fmt.Println("models_UpdateStmt: \n ", models_UpdateStmt)

	// Generate models-UpdateErr

	for i, _ := range UpdateErr {
		//fmt.Println("Valor de i", i, "Valor de j", j)

		multilineAuxUpdateErr = multilineAuxUpdateErr + UpdateErr[i] + "\n"
	}

	models_UpdateErr = "_, err := db.ExecContext(ctx, stmt," + "\n" + multilineAuxUpdateErr + "\t" + "time.Now()," + "\n \t" + "{{.LowerEntity}}.Id," + "\n" + ")"

	fmt.Println("\n")
	fmt.Println("models_UpdateErr: \n ", models_UpdateErr)
	fmt.Println("\n")

	// Generate models-GetAllQuery
	// Finding!!! order  by name depends JSON config file for order by nombre

	models_GetAllQuery = "query := `select id, " + auxGetOneQuery + "created_at, updated_at from {{.LowerEntity}}s order by nombre`"
	// query := `select id, name, description, price, created_at, updated_at from products where id = $1`
	// query := `select id, nombre, descripcion, precio, cantidad, random, created_at, updated_at from {{.LowerEntity}} where id = $1`
	fmt.Println("models_GetAllQuery es: \n ", models_GetAllQuery)
	fmt.Println("\n")

	TypesVars["models-typeEntityStruct"] = models_typeEntityStruct
	TypesVars["models-InsertStmt"] = models_InsertStmt
	TypesVars["models-InsertErr"] = models_InsertErr
	TypesVars["models-GetOneQuery"] = models_GetOneQuery
	TypesVars["models-GetOneErr"] = models_GetOneErr
	TypesVars["models-UpdateStmt"] = models_UpdateStmt
	TypesVars["models-UpdateErr"] = models_UpdateErr
	TypesVars["models-GetAllQuery"] = models_GetAllQuery
	TypesVars["models-GetAllErrRowsScan"] = models_GetAllErrRowsScan

	// GENERANDO TIPOS
	TypesVars["models-DeleteStmt"] = "" // validar si realmente es necesario

	return TypesVars

}

func generateStmtValues(quantity int) string {
	var aux string

	for i := 1; i <= quantity; i++ {
		aux = aux + "$" + strconv.Itoa(i)
		if i != quantity {
			aux = aux + ", "
		}
	}

	return aux
}
