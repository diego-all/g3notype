package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
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
	// query := `Tengo un modelo de datos: ` + class + ` con los siguientes atributos y su tipo de dato correspondiente` + formattedMetadata +
	// 	// nombre | string
	// 	// descripcion| string
	// 	// precio | int
	// 	// cantidad | int
	// 	// random| int
	// 	// created_at|DATETIME('now')
	// 	// updated_at|DATETIME('now')

	// 	`Requiero construir basado en los datos anteriores las sentencias insert con data dummy, en total 5 sentencias para una base de datos sqlite, como las siguientes:

	// 	-- DML statements [Dummy data]
	// 	INSERT INTO products (name, description, price, created_at, updated_at)
	// 	     VALUES ('Teléfono móvil', 'Smartphone de última generación', 799, DATETIME('now'), DATETIME('now'));

	// 	INSERT INTO products (name, description, price, created_at, updated_at)
	// 	     VALUES ('Camiseta', 'Camiseta de algodón', 20, DATETIME('now'), DATETIME('now'));

	// 	INSERT INTO products (name, description, price, created_at, updated_at)
	// 	     VALUES ('Sartén antiadherente', 'Sartén para cocinar', 35, DATETIME('now'), DATETIME('now'));

	// 	INSERT INTO products (name, description, price, created_at, updated_at)
	// 	     VALUES ('Balón de fútbol', 'Balón oficial de la FIFA', 50, DATETIME('now'), DATETIME('now'));

	// 	INSERT INTO products (name, description, price, created_at, updated_at)
	// 	     VALUES ('Muñeca', 'Muñeca de peluche para niños', 15, DATETIME('now'), DATETIME('now'));`

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
			 VALUES ('Muñeca', 'Muñeca de peluche para niños', 15, DATETIME('now'), DATETIME('now'));`

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

	// Unir las partes en una sola cadena de texto
	return fmt.Sprintf("%s", strings.Join(parts, "\n"))
}

func ExtractInsertStatements(data string) string {
	// Utilizar expresión regular para extraer las sentencias INSERT
	re := regexp.MustCompile(`(?i)INSERT INTO [^\;]+;`)
	inserts := re.FindAllString(data, -1)

	// Unir todas las sentencias INSERT en un solo string
	return strings.Join(inserts, "\n")
}

// PENDIENTE RECIBIR ESTE: class string, classMetadata [][]string, AL PARECER CONFIG NO ES GLOBAL.

func AddDummyData(class string, classMetadata [][]string) string {
	// Llamar a GenerateDummyData para obtener los datos dummy

	//config := models.Config{}
	dummyData := GenerateDummyData(class, classMetadata)

	fmt.Println("DESDE ADDDUMMYDATA: (clase) \n", class)

	fmt.Println("\n")

	fmt.Println("DESDE ADDDUMMYDATA: (clase) \n", classMetadata)

	// Extraer solo las sentencias INSERT
	return ExtractInsertStatements(dummyData)
}
