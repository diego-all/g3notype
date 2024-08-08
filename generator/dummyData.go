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

	"github.com/diego-all/run-from-gh/models"
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

func GenerateDummyData(class string, classMetadata [][]string) string {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtener la clave de la API desde las variables de entorno
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

	// AL PARECER NO ES GLOBAL, VALIDAR!!!!
	fmt.Println("CONSULTANDO A GEMINI: (clase) \n", class)

	var formattedMetadata []string

	for _, pair := range classMetadata {
		if len(pair) == 2 {
			formattedMetadata = append(formattedMetadata, fmt.Sprintf("%s|%s", pair[0], pair[1]))
		}
		//fmt.Println("valor de i:", i, "valor de j", j)
	}
	formattedMetadata = append(formattedMetadata, "created_at|DATETIME('now')")
	formattedMetadata = append(formattedMetadata, "updated_at|DATETIME('now')")

	// Unir todas las líneas en un solo string separado por saltos de línea
	formattedMetadataString := strings.Join(formattedMetadata, "\n")

	fmt.Println("FORMATTEDMETADATA: \n", formattedMetadata)

	// Definir la consulta
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
			  
		Es necesario no utilizar caracteres especiales ni comas en los posesivos en la datadummy en caso de ser informacion en ingles.
		Ademas considerar que la entidad para nombrar la tabla debe ser en plural en las sentencias insert.
		
		También requiero que generes a partir de los 2 primeros inserts la estructura de una request JSON. Es decir 2 veces el siguiente
		ejemplo considerando el tipo de dato si son strings utilizar comillas, en caso de ser valores numericos omitirlas.
		Debes omitir agregar los campos created_at y updated_at en estos 2 nuevos strings a generar, tambien omitir las llaves {}, no será
		un JSON, solo es una porcion del mismo.

		"name": "value",
		"description": "value",
		"price": 100000
		`

	fmt.Println("QUERY:\n", query)

	resp, err := model.GenerateContent(
		ctx,
		genai.Text(query),
	)
	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
	}

	// Convertir la respuesta a JSON
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}

	// Deserializar la respuesta JSON en la estructura DDLData
	var data DDLData
	err = json.Unmarshal(respJSON, &data)
	if err != nil {
		log.Fatalf("Error al deserializar la respuesta: %v", err)
	}

	fmt.Println("DATA:\n", data) // No se ve

	// Recopilar el contenido de Parts
	var parts []string
	for _, candidate := range data.Candidates {
		parts = append(parts, candidate.Content.Parts...)
	}

	fmt.Println("PARTS:\n", parts)

	// Unir las partes en una sola cadena de texto
	fmt.Println("GENERATEDUMMYDATA SENTENCIAS INSERT: \n")
	fmt.Println(fmt.Sprintf("%s", strings.Join(parts, "\n")))
	return fmt.Sprintf("%s", strings.Join(parts, "\n"))
}

func ExtractInsertStatements(data string) string {
	// Utilizar expresión regular para extraer las sentencias INSERT
	re := regexp.MustCompile(`(?i)INSERT INTO [^\;]+;`)
	inserts := re.FindAllString(data, -1)

	// Unir todas las sentencias INSERT en un solo string
	return strings.Join(inserts, "\n")
}

// ExtractUpsertCollections extrae dos estructuras JSON del response de Gemini, eliminando los campos created_at y updated_at.
func ExtractUpsertCollections(data string) (string, string) {
	// Regex para encontrar los valores de los primeros dos INSERT
	re := regexp.MustCompile(`(?i)INSERT INTO [^\(]+\(([^\)]+)\)\s+VALUES\s+\(([^\)]+)\);`)
	matches := re.FindAllStringSubmatch(data, -1)

	if len(matches) < 2 {
		return "", ""
	}

	// Convertir los dos primeros inserts a JSON sin los campos created_at y updated_at
	createJSON := convertInsertToJSON(matches[0][2], matches[0][3])
	updateJSON := convertInsertToJSON(matches[1][2], matches[1][3])

	return createJSON, updateJSON
}

// convertInsertToJSON convierte los valores de un INSERT a una estructura JSON
func convertInsertToJSON(columns string, values string) string {
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

	return "{" + strings.Join(result, ", ") + "}"
}

// isNumeric verifica si una cadena representa un valor numérico
func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// PENDIENTE RECIBIR ESTE: class string, classMetadata [][]string, AL PARECER CONFIG NO ES GLOBAL.

// func AddDummyData(class string, classMetadata [][]string) string {
// 	// Llamar a GenerateDummyData para obtener los datos dummy

// 	dummyData := GenerateDummyData(class, classMetadata)

// 	fmt.Println("DESDE ADDDUMMYDATA: (clase) \n", class)

// 	fmt.Println("\n")

// 	fmt.Println("DESDE ADDDUMMYDATA: (clase) \n", classMetadata)

// 	// Extraer solo las sentencias INSERT
// 	return ExtractInsertStatements(dummyData)
// }

func AddDummyData(class string, classMetadata [][]string) models.DummyDataResult {
	// Llamar a GenerateDummyData para obtener los datos dummy
	dummyData := GenerateDummyData(class, classMetadata)
	//insertStatements := ExtractInsertStatements(dummyData)
	createJSON, updateJSON := ExtractUpsertCollections(dummyData)
	//inserts := ExtractInsertStatements(dummyData)

	//fmt.Println("EXTRACTED INSERTS FROM ADDDUMMYDATA(): \n", inserts)

	// var createJSON, updateJSON string
	// if len(inserts) > 0 {
	// 	createJSON = CreateJSONPortionFromInsert(inserts[0])
	// }
	// if len(inserts) > 1 {
	// 	updateJSON = CreateJSONPortionFromInsert(inserts[1])
	// }

	fmt.Println("DESDE ADDDUMMYDATA: (clase) \n", class, classMetadata)
	fmt.Println("\n")

	fmt.Println("ESTRUCTURAS NUEVAS CREATEUPDATE: \n", createJSON, updateJSON)

	//return ExtractInsertStatements(dummyData)
	return models.DummyDataResult{
		//Inserts: insertStatements,
		Inserts: ExtractInsertStatements(dummyData),
		//Inserts: dummyData,
		//Inserts:    strings.Join(inserts, "\n"),
		// CreateJSON: "En construccion",
		// UpdateJSON: "En construccion",
		CreateJSON: createJSON,
		UpdateJSON: updateJSON,
	}
}
