package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
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

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get API key from environment variables
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

	fmt.Println("Consultando dummy data a Gemini:\n", class, "\n")

	var formattedMetadata []string

	for _, pair := range classMetadata {
		if len(pair) == 2 {
			formattedMetadata = append(formattedMetadata, fmt.Sprintf("%s|%s", pair[0], pair[1]))
		}
	}
	formattedMetadata = append(formattedMetadata, "created_at|DATETIME('now')")
	formattedMetadata = append(formattedMetadata, "updated_at|DATETIME('now')")

	// Join all lines into a single string separated by line breaks
	formattedMetadataString := strings.Join(formattedMetadata, "\n")

	//fmt.Println("formattedMetadata: \n", formattedMetadata)

	// Define the query
	query := `Tengo un modelo de datos: ` + class + ` con los siguientes atributos y su tipo de dato correspondiente:
		` + formattedMetadataString + `
				
		Requiero construir basado en los datos anteriores las sentencias insert con data dummy, en total 5 sentencias para una base de datos sqlite, como las siguientes:

		Nota: Debes tomar estos inserts como referencia pero los nombres de los campos son los mencionados anteriormente.
				
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

		Por favor utiliza para el tag createBody para el primero y updateBody para el segundo como se muestra a continuación:

		createBody:
		"name": "value1",
		"description": "value1",
		"price": 100000

		updateBody:
		"name": "value2",
		"description": "value2",
		"price": 1000
		`

	//fmt.Println("Query:\n", query)

	resp, err := model.GenerateContent(
		ctx,
		genai.Text(query),
	)
	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
	}

	// Convert response to JSON
	respJSON, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}

	// Deserialize JSON response into DDLData structure
	var data DDLData
	err = json.Unmarshal(respJSON, &data)
	if err != nil {
		log.Fatalf("Error al deserializar la respuesta: %v", err)
	}

	//fmt.Println("data:\n", data)

	// Collect Parts content
	var parts []string
	for _, candidate := range data.Candidates {
		parts = append(parts, candidate.Content.Parts...)
	}

	//fmt.Println("Parts:\n", parts)

	// Join the parts into a single text string

	//fmt.Println(fmt.Sprintf("%s", strings.Join(parts, "\n")))
	return fmt.Sprintf("%s", strings.Join(parts, "\n"))
}

func ExtractInsertStatements(data string) string {
	// Use regular expression to extract INSERT statements
	re := regexp.MustCompile(`(?i)INSERT INTO [^\;]+;`)
	inserts := re.FindAllString(data, -1)

	// Join all INSERT statements into a single string
	return strings.Join(inserts, "\n")
}

// Function to extract createJSON and updateJSON from two INSERTs
func extractJSONFromInsert(input string) (string, string, error) {
	// Define regular expression to extract fields and values
	pattern := regexp.MustCompile(`INSERT INTO \w+ \((.*?)\)\s*VALUES\s*\((.*?)\);`)

	matches := pattern.FindAllStringSubmatch(input, -1)
	if len(matches) < 2 {
		return "", "", fmt.Errorf("No se encontraron suficientes INSERTs en la entrada")
	}

	// Internal function to build the JSON
	buildJSON := func(fields, values []string) string {
		var jsonBuilder strings.Builder
		for i := 0; i < len(fields); i++ {
			if !strings.Contains(values[i], "DATETIME('now')") {
				field := strings.TrimSpace(fields[i])
				value := strings.TrimSpace(values[i])

				// Add double quotes to string values
				if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
					value = fmt.Sprintf("\"%s\"", strings.Trim(value, "'"))
				}
				jsonBuilder.WriteString(fmt.Sprintf("\"%s\": %s", field, value))
				// Validation to avoid adding comma after last field
				if i < len(fields)-3 {
					jsonBuilder.WriteString(", ")
				}
			}
		}
		return jsonBuilder.String()
	}

	// Get fields and values ​​from first and second INSERT
	fields1 := strings.Split(matches[0][1], ", ")
	values1 := strings.Split(matches[0][2], ", ")
	fields2 := strings.Split(matches[1][1], ", ")
	values2 := strings.Split(matches[1][2], ", ")

	// Build createJSON and updateJSON
	createJSON := buildJSON(fields1, values1)
	updateJSON := buildJSON(fields2, values2)

	return createJSON, updateJSON, nil
}

func AddDummyData(class string, classMetadata [][]string) models.DummyDataResult {
	// Call GenerateDummyData to get the dummy data
	dummyData := GenerateDummyData(class, classMetadata)

	createJSON, updateJSON, _ := extractJSONFromInsert(dummyData)

	//fmt.Println("CREATEJSON \n", createJSON)
	//fmt.Println("UPDATEJSON \n", updateJSON)

	//return ExtractInsertStatements(dummyData)
	return models.DummyDataResult{
		Inserts:    ExtractInsertStatements(dummyData),
		CreateJSON: createJSON,
		UpdateJSON: updateJSON,
	}
}
