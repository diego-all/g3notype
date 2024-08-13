package generator

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
)

var TypesVars = map[string]string{
	// handlers
	"handlers-typeEntityRequest":     "",
	"handlers-typeEntityResponse":    "",
	"handlers-varCreateEntityModels": "",
	"handlers-varGetEntResponse":     "",
	"handlers-varUpdateEntityModels": "",
	"handlers-payloadCreateResponse": "",
	"handlers-payloadUpdateResponse": "",
	//Database_DDL_statement
	"database-DDL-statement": "",
	"database-DummyData":     "",
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
	"models-DeleteStmt":        "", // validate if it is really necessary
	// Requests (Collections)
	"requests-Create": "",
	"requests-Update": "",
}

var preTemplates = map[string]string{
	"cmd/api/handlers-entity-base.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-entity-generic.txt",
	"database/up.sql":                 "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/database/up.sql-generic.txt",
	"internal/entities-base.go":       "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/internal/entities-generic.txt",
	"requests-base.txt":               "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/requests-generic.txt",
	//"cmd/api/handlers-{{.Entity}}.go": "/home/diegoall/MAESTRIA_ING/CLI/run-from-gh/base-templates/cmd/api/handlers-entity-base.txt",
}

var (
	NaturalID string
)

func SetNaturalID(id string) {
	NaturalID = id
}

type PreTemplateData struct {
	Handlers_typeEntityRequest     string
	Handlers_typeEntityResponse    string
	Handlers_varCreateEntityModels string
	Handlers_varGetEntResponse     string
	Handlers_varUpdateEntityModels string
	Handlers_payloadCreateResponse string
	Handlers_payloadUpdateResponse string

	Database_DDL_statement string
	Database_DummyData     string

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

	Collection_Create string
	Collection_Update string

	Entity      string
	LowerEntity string
}

// func generateHandlers(class string, classMetadata map[string]string) (string) {
func generateHandlers(class string, classMetadata [][]string) map[string]string {

	//fmt.Println("Class metadata", classMetadata)
	//longitud := len(classMetadata)
	//fmt.Println("longitud del map es:", longitud)
	//fmt.Println("\n")

	var auxReqRes string
	var auxCreateEntModels string // It was used for Update, only 2 fields are added
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

	var naturalID string

	for k, attribute := range classMetadata {

		attributeName := attribute[0]
		attributeType := attribute[1]

		auxReqRes = "\t" + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + "\t" + attributeType + "\t" + "`json:\"" + attributeName + "\"`"
		auxCreateEntModels = "\t" + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ":\t" + "{{.LowerEntity}}Req." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","
		auxGetEntResponse = "\t" + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ":\t" + "{{.LowerEntity}}." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","

		// Append each of the read attributes
		reqResTypes = append(reqResTypes, auxReqRes)
		createEntModels = append(createEntModels, auxCreateEntModels)
		getEntResponse = append(getEntResponse, auxGetEntResponse)

		if k == 0 {
			SetNaturalID(string(strings.ToUpper(string(attributeName[0])) + string(attributeName[1:])))
		}
	}

	fmt.Println("naturalID for model: \n ", naturalID)

	// They are verticalized, I think it look better with a while
	for i, _ := range reqResTypes {
		//fmt.Println("Valor de i", i, "Valor de j", j)
		multilineAuxReqResTypes = multilineAuxReqResTypes + reqResTypes[i] + "\n"
	}

	handlers_typeEntityRequest = "type {{.LowerEntity}}Request struct {" + "\n" + multilineAuxReqResTypes + "}"
	//fmt.Println("handlers_typeEntityRequest: \n ", handlers_typeEntityRequest)

	handlers_typeEntityResponse = "type {{.LowerEntity}}Response struct {" + "\n" + multilineAuxReqResTypes + "}"
	//fmt.Println("handlers_typeEntityResponse: \n ", handlers_typeEntityResponse)

	// For createEntModels
	for i, _ := range createEntModels {
		multilineAuxCEntModels = multilineAuxCEntModels + createEntModels[i] + "\n"
	}

	// For create handlers_varCreateEntityModels
	handlers_varCreateEntityModels = "var {{.LowerEntity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "}"
	//fmt.Println("handlers_varCreateEntityModels: \n ", handlers_varCreateEntityModels)

	// Para update handlers_varUpdateEntResponse
	handlers_varUpdateEntityModels = "var {{.LowerEntity}} = models.{{.Entity}}{" + "\n" + multilineAuxCEntModels + "\t" + "UpdatedAt:   time.Now()," + "\n \t" + "Id:          {{.LowerEntity}}ID," + "\n" + "}"
	//fmt.Println("handlers_varUpdateEntityModels: \n ", handlers_varUpdateEntityModels)

	for i, _ := range getEntResponse {
		multilineAuxGEntResponse = multilineAuxGEntResponse + getEntResponse[i] + "\n"
	}

	handlers_varGetEntResponse = "var {{.LowerEntity}}Response = {{.LowerEntity}}Response{\n" + multilineAuxGEntResponse + "}"
	//fmt.Println("handlers_varGetEntResponse: \n ", handlers_varGetEntResponse)

	handlers_payloadCreateResponse = "payload = jsonResponse{\n" + "\t    Error:   false,\n" + "\t    Message: \"{{.Entity}} successfully created\",\n" + "\t    Data:    envelope{\"{{.LowerEntity}}\": {{.LowerEntity}}." + NaturalID + "},\n" + "\t}"
	//fmt.Println("handlers_payloadCreateResponse: \n ", handlers_payloadCreateResponse)

	handlers_payloadUpdateResponse = "payload = jsonResponse{\n" + "\t    Error:   false,\n" + "\t    Message: \"{{.Entity}} successfully updated\",\n" + "\t    Data:    envelope{\"{{.LowerEntity}}\": {{.LowerEntity}}." + NaturalID + "},\n" + "\t}"
	//fmt.Println("handlers_payloadUpdateResponse: \n ", handlers_payloadUpdateResponse)

	// Generated Types and Vars
	TypesVars["handlers-typeEntityRequest"] = handlers_typeEntityRequest
	TypesVars["handlers-typeEntityResponse"] = handlers_typeEntityResponse
	TypesVars["handlers-varCreateEntityModels"] = handlers_varCreateEntityModels
	TypesVars["handlers-varGetEntResponse"] = handlers_varGetEntResponse
	TypesVars["handlers-varUpdateEntityModels"] = handlers_varUpdateEntityModels
	TypesVars["handlers-payloadCreateResponse"] = handlers_payloadCreateResponse
	TypesVars["handlers-payloadUpdateResponse"] = handlers_payloadUpdateResponse

	return TypesVars
}

func modifyBaseTemplates(preGeneratedTypes map[string]string) {

	//fmt.Println("PREGENERATED TYPES en modifyBaseTemplates: \n", preGeneratedTypes)
	//fmt.Println("LONGITUD DE PREREGENERATED TYPES en modifyBaseTemplates", len(preGeneratedTypes))

	count := 0
	for _, value := range preGeneratedTypes {
		// Check if the value is not empty. Adjusts the condition based on the type of expected value.
		if value != "" { // This is just for string values.
			count++
		}
	}
	//fmt.Println("Número de keys llenas en preGeneratedTypes:", count)

	preData := PreTemplateData{
		Handlers_typeEntityRequest:     preGeneratedTypes["handlers-typeEntityRequest"],
		Handlers_typeEntityResponse:    preGeneratedTypes["handlers-typeEntityResponse"],
		Handlers_varCreateEntityModels: preGeneratedTypes["handlers-varCreateEntityModels"],
		Handlers_varGetEntResponse:     preGeneratedTypes["handlers-varGetEntResponse"],
		Handlers_varUpdateEntityModels: preGeneratedTypes["handlers-varUpdateEntityModels"],
		Handlers_payloadCreateResponse: preGeneratedTypes["handlers-payloadCreateResponse"],
		Handlers_payloadUpdateResponse: preGeneratedTypes["handlers-payloadUpdateResponse"],

		Database_DDL_statement: preGeneratedTypes["database-DDL-statement"],
		Database_DummyData:     preGeneratedTypes["database-DummyData"],

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

		Collection_Create: preGeneratedTypes["requests-Create"],
		Collection_Update: preGeneratedTypes["requests-Update"],

		Entity: "{{.Entity}}",
		//entity:  "{{.entity}}",   // NO funciona con minusculas seguir indagando
		LowerEntity: "{{.LowerEntity}}",
	}

	for projectFile, templatePath := range preTemplates {

		//fmt.Println("Path y Content es: ", projectFile, templatePath)

		// If there is template content, process it
		if templatePath != "" {

			content, err := ioutil.ReadFile(templatePath)
			if err != nil {
				fmt.Println("Error al leer la plantilla:", err)
				continue
			}

			tmpl, err := template.New("fileContent").Parse(string(content))
			//Important for debugging
			//fmt.Println("tmpl es:", tmpl)
			if err != nil {
				fmt.Println("Error al parsear la plantilla:", err)
				continue
			}

			// Create or open the target file in write mode
			file, err := os.Create(templatePath)
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
	fmt.Print("\n")
	fmt.Println("Procesando templates genéricas ... \n")

}

// Generate Create table
// func generateDatabaseDDL(class string, classMetadata [][]string, dummy bool) map[string]string {
func generateDatabaseDDL(class string, classMetadata [][]string, dummy bool) {

	//fmt.Println("Desde generateDDLStatement", class)

	//fmt.Println("Class metadata", classMetadata)
	//longitud := len(classMetadata)
	//fmt.Println("longitud del map es:", longitud)

	var auxDDL, auxCollection string
	var ddlStatement, collectionRequest []string

	var sqliteValue string
	var multilineAuxDDLStatement string

	var Database_DDL_statement, Database_DummyData, Collection_Create, Collection_Update string

	for i, data := range classMetadata {

		attributeName := data[0]
		attributeType := data[1]

		switch attributeType {
		case "int":
			sqliteValue = "INTEGER"
		case "string":
			sqliteValue = "VARCHAR(100)"
		case "bool":
			sqliteValue = "INTEGER"
		case "":
			fmt.Println("OTRO CASO")

		}

		auxDDL = "\t" + attributeName + " " + sqliteValue + " " + "NOT NULL,"
		ddlStatement = append(ddlStatement, auxDDL)

		// Check if it is the last iteration to avoid adding the comma
		if i == len(classMetadata)-1 {
			if sqliteValue == "VARCHAR(100)" {
				auxCollection = "\"" + attributeName + "\": " + "\"value\""
			} else {
				auxCollection = "\"" + attributeName + "\": " + "10"
			}
		} else {
			if sqliteValue == "VARCHAR(100)" {
				auxCollection = "\"" + attributeName + "\": " + "\"value\","
			} else {
				auxCollection = "\"" + attributeName + "\": " + "10,"
			}
		}

		collectionRequest = append(collectionRequest, auxCollection)

	}

	// They are verticalized, I think it look better with a while
	for i, _ := range ddlStatement {
		multilineAuxDDLStatement = multilineAuxDDLStatement + ddlStatement[i] + "\n"
	}

	Database_DDL_statement = "CREATE TABLE IF NOT EXISTS {{.LowerEntity}}s (\n \t" + "id INTEGER PRIMARY KEY AUTOINCREMENT,\n" + multilineAuxDDLStatement + "\t" + "created_at TIMESTAMP DEFAULT DATETIME NOT NULL,\n \t" + "updated_at TIMESTAMP NOT NULL\n \t" + ");"

	Database_DummyData = ""

	for i, _ := range collectionRequest {
		Collection_Create = Collection_Create + collectionRequest[i] + "\n"
	}

	//fmt.Println("Dummy: \n", dummy)

	if dummy {

		dummyDataResult := AddDummyData(class, classMetadata)

		Database_DummyData = dummyDataResult.Inserts
		Collection_Create = dummyDataResult.CreateJSON
		Collection_Update = dummyDataResult.UpdateJSON

		// fmt.Println("El valor de Database_DummyData es:\n", dummyDataResult.Inserts)
		// fmt.Println("El valor del JSON para CREATE es:\n", dummyDataResult.CreateJSON)
		// fmt.Println("El valor del JSON para UPDATE es:\n", dummyDataResult.UpdateJSON)
	}

	TypesVars["database-DDL-statement"] = Database_DDL_statement
	TypesVars["database-DummyData"] = Database_DummyData

	TypesVars["requests-Create"] = Collection_Create

	TypesVars["requests-Update"] = Collection_Update

	TypesVars["requests-Create"] = Collection_Create

	TypesVars["requests-Update"] = Collection_Create

	TypesVars["requests-Update"] = Collection_Update

	//return TypesVars
}

// Generate
func generateEntityModels(class string, classMetadata [][]string) map[string]string {

	//fmt.Println("Class metadata", classMetadata)
	longitud := len(classMetadata)
	//fmt.Println("longitud del map es:", longitud)

	// Generate models-typeEntityStruct
	// Attention with "name,omitempty"`
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

		auxTypeEntityStruct = strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + "\t" + attributeType + "\t" + "`json:\"" + attributeName + "\"`"
		auxInsertStmt = auxInsertStmt + attributeName + ", "
		auxInsertErr = "\t" + "{{.LowerEntity}}." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","
		auxGetOneQuery = auxGetOneQuery + attributeName + ", "
		auxGetOneErr = "\t" + "&{{.LowerEntity}}." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","
		auxUpdateStmt = "\t" + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + " = $" + strconv.Itoa(i) + ","
		auxUpdateErr = "\t" + "{{.LowerEntity}}." + strings.ToUpper(string(attributeName[0])) + string(attributeName[1:]) + ","

		typeEntityStructs = append(typeEntityStructs, auxTypeEntityStruct)
		InsertErrs = append(InsertErrs, auxInsertErr)
		GetOneErrs = append(GetOneErrs, auxGetOneErr)
		UpdateStmt = append(UpdateStmt, auxUpdateStmt)
		UpdateErr = append(UpdateErr, auxUpdateErr)
		i++
	}

	// They are verticalized, I think it look better with a while
	for i, _ := range typeEntityStructs {
		multilineAuxTypeEntityStructs = multilineAuxTypeEntityStructs + "\t" + typeEntityStructs[i] + "\n"
	}

	models_typeEntityStruct = "type {{.Entity}} struct {" + "\n \t" + "Id	int	`json:\"id\"`" + "\n" + multilineAuxTypeEntityStructs + "\t" + "CreatedAt   time.Time `json:\"created_at\"`" + "\n \t" + "UpdatedAt   time.Time `json:\"updated_at\"`" + "\n" + "}"
	//fmt.Println("models_typeEntityStruct: \n ", models_typeEntityStruct)

	// Generate models-InsertStmt
	generateStmtValues(longitud + 2) // created_at, updated_at
	models_InsertStmt = "\tstmt := `insert into {{.LowerEntity}}s (" + auxInsertStmt + "created_at, updated_at)\n \t" + "values (" + generateStmtValues(longitud+2) + ")" + " returning  id`"
	//fmt.Println("models_InsertStmt es: \n", models_InsertStmt)

	// Generate models-InsertErr
	for i, _ := range InsertErrs {
		multilineAuxInsertErr = multilineAuxInsertErr + InsertErrs[i] + "\n"
	}

	models_InsertErr = "err := db.QueryRowContext(ctx, stmt," + "\n" + multilineAuxInsertErr + "\t" + "time.Now()," + "\n" + "\t" + "time.Now()," + "\n" + ").Scan(&newID)"

	//fmt.Println("models_InsertErr es: \n ", models_InsertErr)

	// Generate models-GetOneQuery
	models_GetOneQuery = "\tquery := `select id, " + auxGetOneQuery + "created_at, updated_at from {{.LowerEntity}}s where id = $1`"
	//fmt.Println("models_GetOneQuery es: \n ", models_GetOneQuery)

	// Generate models-GetOneErr
	for i, _ := range GetOneErrs {
		multilineAuxGetOneErr = multilineAuxGetOneErr + GetOneErrs[i] + "\n"
	}

	models_GetOneErr = "err := row.Scan(" + "\n" + "\t" + "&{{.LowerEntity}}.Id," + multilineAuxGetOneErr + "\n \t" + "&{{.LowerEntity}}.CreatedAt," + "\n" + "\t" + "&{{.LowerEntity}}.UpdatedAt," + "\n" + ")"
	//fmt.Println("models_GetOneErr es: \n ", models_GetOneErr)

	// Generate models-GetAllErrRowsScan
	models_GetAllErrRowsScan = "err := rows.Scan(" + "\n" + "\t" + "&{{.LowerEntity}}.Id," + multilineAuxGetOneErr + "\n \t" + "&{{.LowerEntity}}.CreatedAt," + "\n" + "\t" + "&{{.LowerEntity}}.UpdatedAt," + "\n" + ")"
	//fmt.Println("models_GetAllErrRowsScan es: \n ", models_GetAllErrRowsScan)

	// Generate models-UpdateStmt
	for i, _ := range UpdateStmt {
		multilineAuxUpdateStmt = multilineAuxUpdateStmt + UpdateStmt[i] + "\n"
	}

	models_UpdateStmt = "stmt := `update {{.LowerEntity}}s set" + "\n" + " " + multilineAuxUpdateStmt + "\t" + "updated_at = $" + strconv.Itoa(i+1) + "\n \t" + "where id = $" + strconv.Itoa(i+2) + "`"
	//fmt.Println("models_UpdateStmt: \n ", models_UpdateStmt)

	// Generate models-UpdateErr
	for i, _ := range UpdateErr {
		multilineAuxUpdateErr = multilineAuxUpdateErr + UpdateErr[i] + "\n"
	}

	models_UpdateErr = "_, err := db.ExecContext(ctx, stmt," + "\n" + multilineAuxUpdateErr + "\t" + "time.Now()," + "\n \t" + "{{.LowerEntity}}.Id," + "\n" + ")"
	//fmt.Println("models_UpdateErr: \n ", models_UpdateErr)

	// Generate models-GetAllQuery
	models_GetAllQuery = "query := `select id, " + auxGetOneQuery + "created_at, updated_at from {{.LowerEntity}}s order by " + strings.ToLower(NaturalID) + "`"

	TypesVars["models-typeEntityStruct"] = models_typeEntityStruct
	TypesVars["models-InsertStmt"] = models_InsertStmt
	TypesVars["models-InsertErr"] = models_InsertErr
	TypesVars["models-GetOneQuery"] = models_GetOneQuery
	TypesVars["models-GetOneErr"] = models_GetOneErr
	TypesVars["models-UpdateStmt"] = models_UpdateStmt
	TypesVars["models-UpdateErr"] = models_UpdateErr
	TypesVars["models-GetAllQuery"] = models_GetAllQuery
	TypesVars["models-GetAllErrRowsScan"] = models_GetAllErrRowsScan
	TypesVars["models-DeleteStmt"] = ""

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
