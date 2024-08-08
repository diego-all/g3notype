Tengo un script en golang que realiza una consulta a gemini y obtiene el sigueinte response:

```sql
-- DML statements [Dummy data]
INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
 VALUES ('Telefono movil', 'Smartphone de ultima generacion', 799, 12, 'modelo1', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
 VALUES ('Camiseta', 'Camiseta de algodon', 20, 24, 'modelo2', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
 VALUES ('Sarten antiadherente', 'Sarten para cocinar', 35, 36, 'modelo3', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
 VALUES ('Balon de futbol', 'Balon oficial de la FIFA', 50, 48, 'modelo4', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
 VALUES ('Muneca', 'Muneca de peluche para ninos', 15, 60, 'modelo5', DATETIME('now'), DATETIME('now'));
```

## Estructura JSON para los dos primeros inserts

```json
"nombresito": "Telefono movil",
"descripcionsita": "Smartphone de ultima generacion",
"precioaquel": 799,
"cantidadparce": 12,
"randomoelo": "modelo1"

"nombresito": "Camiseta",
"descripcionsita": "Camiseta de algodon",
"precioaquel": 20,
"cantidadparce": 24,
"randomoelo": "modelo2"
``` 

Este response seria data en GenerateDummyData()

ya tengo la funcion con la logica para extraer los 5 Inserts y asignarlos a una variable de tipo string.
La funcion seria ExtractInsertStatements()

Requiero construir una logica similar, puede ser con regex u de otra forma con el fin de extraer las 2 estructuras siguientes que utilizan la misma data dumy de los 2 primeros inserts, son porciones de datadummy para llenar el body de un JSON.

A continuacion se encuentra el script:

package models

type DummyDataResult struct {
	Inserts    string
	CreateJSON string
	UpdateJSON string
}

Esta es una porcion de la funcion que invoca el script:

package generator

func generateDatabaseDDL(class string, classMetadata [][]string, dummy bool) map[string]string {

	if dummy {

		dummyDataResult := AddDummyData(class, classMetadata)

		Database_DummyData = dummyDataResult.Inserts
		Collection_Create = dummyDataResult.CreateJSON
		Collection_Update = dummyDataResult.UpdateJSON

		fmt.Println("El valor de Database_DummyData es:\n", dummyDataResult.Inserts)
		fmt.Println("El valor del JSON para CREATE es:\n", dummyDataResult.CreateJSON)
		fmt.Println("El valor del JSON para UPDATE es:\n", dummyDataResult.UpdateJSON)
	}

}


package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type DDLData struct {
	Candidates []struct {
		Index   int `json:"Index"`
		Content struct {
			Parts []string `json:"Parts"`
			Role  string   `json:"Role"`
		} `json:"Content"`
		FinishReason  int `json:"FinishReason"`
		SafetyRatings []struct {
			Category    int  `json:"Category"`
			Probability int  `json:"Probability"`
			Blocked     bool `json:"Blocked"`
		} `json:"SafetyRatings"`
		CitationMetadata interface{} `json:"CitationMetadata"`
		TokenCount       int         `json:"TokenCount"`
	} `json:"Candidates"`
	PromptFeedback interface{} `json:"PromptFeedback"`
	UsageMetadata  struct {
		PromptTokenCount        int `json:"PromptTokenCount"`
		CachedContentTokenCount int `json:"CachedContentTokenCount"`
		CandidatesTokenCount    int `json:"CandidatesTokenCount"`
		TotalTokenCount         int `json:"TotalTokenCount"`
	} `json:"UsageMetadata"`
}

type DummyDataResult struct {
	Inserts    string `json:"inserts"`
	CreateJSON string `json:"create_json"`
	UpdateJSON string `json:"update_json"`
}

func GenerateDummyData(class string, classMetadata [][]string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatalf("GEMINI_API_KEY not found in environment variables")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	model := client.GenerativeModel("gemini-1.5-flash")

	var formattedMetadata []string
	for _, pair := range classMetadata {
		if len(pair) == 2 {
			formattedMetadata = append(formattedMetadata, fmt.Sprintf("%s|%s", pair[0], pair[1]))
		}
	}
	formattedMetadata = append(formattedMetadata, "created_at|DATETIME('now')")
	formattedMetadata = append(formattedMetadata, "updated_at|DATETIME('now')")

	formattedMetadataString := strings.Join(formattedMetadata, "\n")

	query := `Tengo un modelo de datos: ` + class + ` con los siguientes atributos y su tipo de dato correspondiente:
		` + formattedMetadataString + `
		Requiero construir basado en los datos anteriores las sentencias insert con data dummy, en total 5 sentencias para una base de datos sqlite, como las siguientes:
		-- DML statements [Dummy data]
		INSERT INTO products (name, description, price, created_at, updated_at)
			 VALUES ('Teléfono móvil', 'Smartphone de última generación', 799, DATETIME('now'), DATETIME('now'));
		INSERT INTO products (name, description, price, created_at, updated_at)
			 VALUES ('Camiseta', 'Camiseta de algodón', 20, DATETIME('now'), DATETIME('now'));
		INSERT INTO products (name, description, price, created_at, updated_at)
			 VALUES ('Sartén antiadherente', 'Sartén para cocinar', 35, DATETIME('now'), DATETIME('now'));
		INSERT INTO products (name, description, price, created_at, updated_at)
			 VALUES ('Balón de fútbol', 'Balón oficial de la FIFA', 50, DATETIME('now'), DATETIME('now'));
		INSERT INTO products (name, description, price, created_at, updated_at)
			 VALUES ('Muñeca', 'Muñeca de peluche para niños', 15, DATETIME('now'), DATETIME('now'));
		Es necesario no utilizar caracteres especiales ni comas en los posesivos en caso de ser información en inglés.
		Además considerar que la entidad para nombrar la tabla debe ser en plural en las sentencias insert.
		
		También requiero que generes a partir de los 2 primeros inserts la estructura de una request JSON. Es decir, 2 veces el siguiente ejemplo considerando el tipo de dato si son strings utilizar comillas, en caso de ser valores numéricos omitirlas.
		"name": "value",
		"description": "value",
		"price": 100000
		`

	resp, err := model.GenerateContent(
		ctx,
		genai.Text(query),
	)
	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}

	var data DDLData
	err = json.Unmarshal(respJSON, &data)
	if err != nil {
		log.Fatalf("Error al deserializar la respuesta: %v", err)
	}

	var parts []string
	for _, candidate := range data.Candidates {
		parts = append(parts, candidate.Content.Parts...)
	}

	return strings.Join(parts, "\n")
}

func ExtractInsertStatements(data string) string {
	re := regexp.MustCompile(`(?i)INSERT INTO [^\;]+;`)
	inserts := re.FindAllString(data, -1)
	return strings.Join(inserts, "\n")
}

func ExtractCreateUpdate(data string) (string, string) {
	re := regexp.MustCompile(`(?i)INSERT INTO ([^\(]+)\(([^\)]+)\)\s+VALUES\s+\(([^\)]+)\);`)
	matches := re.FindAllStringSubmatch(data, -1)

	if len(matches) < 2 {
		return "", ""
	}

	create := convertInsertToCurlBody(matches[0][2], matches[0][3])
	update := convertInsertToCurlBody(matches[1][2], matches[1][3])

	return create, update
}

func convertInsertToCurlBody(columns string, values string) string {
	columnList := strings.Split(columns, ",")
	valueList := strings.Split(values, ",")

	var result []string
	for i := range columnList {
		column := strings.TrimSpace(columnList[i])
		value := strings.TrimSpace(valueList[i])
		if column != "created_at" && column != "updated_at" {
			if isNumeric(value) {
				result = append(result, fmt.Sprintf("\"%s\": %s", column, value))
			} else {
				result = append(result, fmt.Sprintf("\"%s\": \"%s\"", column, value))
			}
		}
	}

	return strings.Join(result, ",\n")
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func AddDummyData(class string, classMetadata [][]string) DummyDataResult {
	dummyData := GenerateDummyData(class, classMetadata)
	insertStatements := ExtractInsertStatements(dummyData)
	createJSON, updateJSON := ExtractCreateUpdate(dummyData)

	return DummyDataResult{
		Inserts:    insertStatements,
		CreateJSON: "en construccion",
		UpdateJSON: "en construccion",
	}
}


Por favor no modifiques nada de la logica existente, ni la consulta a gemini (query).
Cabe aclarar que los 5 INSERT estan bien como estan considerando los campos created_at, updated_at.

Sin embargo, para los 2 strings nuevos que debes generar la logica de extraccion no se necesitan los campos created_at ni updated_at.
Por favor podrias crear una funcion llamada ExtractUpsertCollections() que extraiga las dos estructuras del response que entrega gemini:

```json
"nombresito": "Telefono movil",
"descripcionsita": "Smartphone de ultima generacion",
"precioaquel": 799,
"cantidadparce": 12,
"randomoelo": "modelo1"

"nombresito": "Camiseta",
"descripcionsita": "Camiseta de algodon",
"precioaquel": 20,
"cantidadparce": 24,
"randomoelo": "modelo2"
``` 
Y las retorne en 2 strings que sean asignadas a esta estructura analoga a la funcion que extrae los 5 inserts.

return DummyDataResult{
	Inserts:    insertStatements,
	CreateJSON: "en construccion",
	UpdateJSON: "en construccion",
}

