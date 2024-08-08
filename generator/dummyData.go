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

// func extractCreateJSON(input string) (string, error) {
// 	// Definir expresión regular para extraer el bloque createJSON
// 	createPattern := regexp.MustCompile(`VALUES \((.*?)\);`)
// 	match := createPattern.FindStringSubmatch(input)
// 	if len(match) < 2 {
// 		return "", fmt.Errorf("No se encontró createJSON en la entrada")
// 	}
// 	// Remover DATETIME('now') del createJSON
// 	createJSON := match[1]
// 	createJSON = regexp.MustCompile(`,\s*DATETIME\('now'\)`).ReplaceAllString(createJSON, "")
// 	return createJSON, nil
// }

// func extractUpdateJSON(input string) (string, error) {
// 	// Definir expresión regular para extraer el bloque updateJSON
// 	updatePattern := regexp.MustCompile(`VALUES \((.*?)\);`)
// 	match := updatePattern.FindStringSubmatch(input)
// 	if len(match) < 2 {
// 		return "", fmt.Errorf("No se encontró updateJSON en la entrada")
// 	}
// 	// Remover DATETIME('now') del updateJSON
// 	updateJSON := match[1]
// 	updateJSON = regexp.MustCompile(`,\s*DATETIME\('now'\)`).ReplaceAllString(updateJSON, "")
// 	return updateJSON, nil
// }

// Función para extraer createJSON
// func extractCreateJSON(input string) (string, error) {
// 	// Definir expresión regular para extraer los campos y valores
// 	pattern := regexp.MustCompile(`INSERT INTO \w+ \((.*?)\)\s*VALUES\s*\((.*?)\);`)
// 	match := pattern.FindStringSubmatch(input)
// 	if len(match) < 3 {
// 		return "", fmt.Errorf("No se encontró createJSON en la entrada")
// 	}

// 	// Extraer nombres de los campos y valores correspondientes
// 	fields := strings.Split(match[1], ", ")
// 	values := strings.Split(match[2], ", ")

// 	// Construir el string createJSON
// 	var createJSONBuilder strings.Builder
// 	for i, field := range fields {
// 		if !strings.Contains(values[i], "DATETIME('now')") {
// 			field = strings.TrimSpace(field)
// 			value := strings.TrimSpace(values[i])

// 			// Agregar comillas dobles a los valores de tipo string
// 			if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
// 				value = fmt.Sprintf("\"%s\"", strings.Trim(value, "'"))
// 			}
// 			createJSONBuilder.WriteString(fmt.Sprintf("\"%s\": %s", field, value))
// 			if i < len(fields)-1 {
// 				createJSONBuilder.WriteString(", ")
// 			}
// 		}
// 	}
// 	return createJSONBuilder.String(), nil
// }

// func extractUpdateJSON(input string) (string, error) {
// 	// Definir expresión regular para extraer los campos y valores
// 	pattern := regexp.MustCompile(`INSERT INTO \w+ \((.*?)\)\s*VALUES\s*\((.*?)\);`)
// 	match := pattern.FindStringSubmatch(input)
// 	if len(match) < 3 {
// 		return "", fmt.Errorf("No se encontró updateJSON en la entrada")
// 	}

// 	// Extraer nombres de los campos y valores correspondientes
// 	fields := strings.Split(match[1], ", ")
// 	values := strings.Split(match[2], ", ")

// 	// Construir el string updateJSON
// 	var updateJSONBuilder strings.Builder
// 	for i, field := range fields {
// 		if !strings.Contains(values[i], "DATETIME('now')") {
// 			field = strings.TrimSpace(field)
// 			value := strings.TrimSpace(values[i])

// 			// Agregar comillas dobles a los valores de tipo string
// 			if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
// 				value = fmt.Sprintf("\"%s\"", strings.Trim(value, "'"))
// 			}
// 			updateJSONBuilder.WriteString(fmt.Sprintf("\"%s\": %s", field, value))
// 			if i < len(fields)-1 {
// 				updateJSONBuilder.WriteString(", ")
// 			}
// 		}
// 	}
// 	return updateJSONBuilder.String(), nil
// }

// Función para extraer createJSON y updateJSON a partir de dos INSERTs
func extractJSONFromInsert(input string) (string, string, error) {
	// Definir expresión regular para extraer los campos y valores
	pattern := regexp.MustCompile(`INSERT INTO \w+ \((.*?)\)\s*VALUES\s*\((.*?)\);`)
	matches := pattern.FindAllStringSubmatch(input, -1)
	if len(matches) < 2 {
		return "", "", fmt.Errorf("No se encontraron suficientes INSERTs en la entrada")
	}

	// Función interna para construir el JSON
	buildJSON := func(fields, values []string) string {
		var jsonBuilder strings.Builder
		for i := 0; i < len(fields); i++ {
			if !strings.Contains(values[i], "DATETIME('now')") {
				field := strings.TrimSpace(fields[i])
				value := strings.TrimSpace(values[i])

				// Agregar comillas dobles a los valores de tipo string
				if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
					value = fmt.Sprintf("\"%s\"", strings.Trim(value, "'"))
				}
				jsonBuilder.WriteString(fmt.Sprintf("\"%s\": %s", field, value))
				if i < len(fields)-1 { // Evitar agregar coma después del último campo
					jsonBuilder.WriteString(", ")
				}
			}
		}
		return jsonBuilder.String()
	}

	// Obtener campos y valores del primer y segundo INSERT
	fields1 := strings.Split(matches[0][1], ", ")
	values1 := strings.Split(matches[0][2], ", ")
	fields2 := strings.Split(matches[1][1], ", ")
	values2 := strings.Split(matches[1][2], ", ")

	// Construir createJSON y updateJSON
	createJSON := buildJSON(fields1, values1)
	updateJSON := buildJSON(fields2, values2)

	return createJSON, updateJSON, nil
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

	// createJSON, _ := extractCreateJSON(dummyData)

	// updateJSON, _ := extractUpdateJSON(dummyData)

	createJSON, updateJSON, _ := extractJSONFromInsert(dummyData)

	fmt.Println("DUMMYDATA DUMMYDATA DUMMYDATA DUMMYDATA DUMMYDATA: \n", dummyData)
	fmt.Println("\n")

	// fmt.Println("CREATEJSON CREATEJSON CREATEJSON:", createJSON)

	fmt.Println("updateJSON:")
	//insertStatements := ExtractInsertStatements(dummyData)
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

	//fmt.Println("ESTRUCTURAS NUEVAS CREATEUPDATE: \n", createJSON, updateJSON)

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
