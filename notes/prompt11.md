Podrias ayudarme a corregir un error en un script que generaste al parecer esta haciendo referencia a un indice que no existe:
El error esta en la funcion CreateJSONFromInsert(insert string) string. invocada por esta parte del codigo:

	// GEMINI
	if dummy {

		Database_DummyData = AddDummyData(class, classMetadata)

		fmt.Println("EL VALOR DE DE DATABASE_DUMMYDATA ES:\n", Database_DummyData)
	}

DUMMY ES: 
 true
panic: runtime error: index out of range [6] with length 6

goroutine 1 [running]:
github.com/diego-all/run-from-gh/generator.CreateJSONFromInsert({0xc000212058, 0xec})
        /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/generator/dummyData.go:159 +0x339
github.com/diego-all/run-from-gh/generator.AddDummyData({0xc0000a6230?, 0xc000080028?}, {0xc000000780?, 0x2?, 0x2?})
        /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/generator/dummyData.go:182 +0x7a
github.com/diego-all/run-from-gh/generator.generateDatabaseDDL({0xc0000a6230, 0x7}, {0xc000000780, 0x5, 0x8}, 0x1)
        /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/generator/preTemplates.go:492 +0xaa6
github.com/diego-all/run-from-gh/generator.Generate({0x7ffd427c55e4, 0xb}, {0x7ffd427c55b8, 0x6}, {{0x7ffd427c55e4, 0xb}, {0xc0000a6230, 0x7}, {0xc000000780, 0x5, ...}}, ...)
        /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/generator/generator.go:37 +0x37b
github.com/diego-all/run-from-gh/cmd.init.func1(0xc00018e500?, {0xc0000aaf60, 0x1, 0xc3a389?})
        /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/cmd/init.go:64 +0x465
github.com/spf13/cobra.(*Command).execute(0x129b9e0, {0xc0000aaf00, 0x6, 0x6})
        /home/diegoall/go/pkg/mod/github.com/spf13/cobra@v1.8.1/command.go:989 +0xab1
github.com/spf13/cobra.(*Command).ExecuteC(0x129bfa0)
        /home/diegoall/go/pkg/mod/github.com/spf13/cobra@v1.8.1/command.go:1117 +0x3ff
github.com/spf13/cobra.(*Command).Execute(...)
        /home/diegoall/go/pkg/mod/github.com/spf13/cobra@v1.8.1/command.go:1041
github.com/diego-all/run-from-gh/cmd.Execute()
        /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/cmd/root.go:15 +0x1a
main.main()
        /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/main.go:7 +0xf
exit status 2

Basicamente el script realiza una consulta a gemini con el fin de generar data dummy para crear 5 inserts para una base de datos sqlite asociados a un modelo de datos que se le suministra a la consulta (query).

El objetivo de la nueva feature es reutilizar los 2 primeros inserts generados para llenar la nueva structura base del body de 2 JSON que seran utilizados en unas request. (Se aclara que deben ser diferentes por que uno es para una operacion de CREATE y el otro para UPDATE).
Cabe aclarar que no se requieren los campos created_at y update_at en estos 2 nuevos strings a generar. La parte de los 5 inserts no debe ser modificada, ya que esta correcta y ha sido probada.


Aca esta el script:

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

// DummyDataResult estructura para agrupar los resultados
type DummyDataResult struct {
	Inserts    string
	CreateJSON string
	UpdateJSON string
}

// Estructura de respuesta de Gemini

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

// Generar datos dummy usando Gemini
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
			  
		También requiero que generes a partir de los 2 primeros inserts la estructura de una request JSON. Es decir 2 veces el siguiente
		ejemplo considerando el tipo de dato si son strings utilizar comillas, en caso de ser valores numéricos omitirlas.

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

// Extraer las sentencias INSERT del resultado
func ExtractInsertStatements(data string) []string {
	re := regexp.MustCompile(`(?i)INSERT INTO [^\;]+;`)
	inserts := re.FindAllString(data, -1)
	return inserts
}

// Crear una estructura JSON a partir de un insert
func CreateJSONFromInsert(insert string) string {
	re := regexp.MustCompile(`\(([^\)]+)\)\s+VALUES\s+\(([^\)]+)\)`)
	matches := re.FindStringSubmatch(insert)
	if len(matches) < 3 {
		return ""
	}

	fields := strings.Split(matches[1], ", ")
	values := strings.Split(matches[2], ", ")

	jsonMap := make(map[string]string)
	for i, field := range fields {
		value := strings.Trim(values[i], "'")
		if value != "null" {
			if _, err := strconv.Atoi(value); err == nil {
				jsonMap[field] = value
			} else {
				jsonMap[field] = fmt.Sprintf("\"%s\"", value)
			}
		} else {
			jsonMap[field] = value
		}
	}

	jsonData, _ := json.Marshal(jsonMap)
	return string(jsonData)
}

// Función para agregar datos dummy
func AddDummyData(class string, classMetadata [][]string) DummyDataResult {
	dummyData := GenerateDummyData(class, classMetadata)
	inserts := ExtractInsertStatements(dummyData)

	var createJSON, updateJSON string
	if len(inserts) > 0 {
		createJSON = CreateJSONFromInsert(inserts[0])
	}
	if len(inserts) > 1 {
		updateJSON = CreateJSONFromInsert(inserts[1])
	}

	return DummyDataResult{
		Inserts:    strings.Join(inserts, "\n"),
		CreateJSON: createJSON,
		UpdateJSON: updateJSON,
	}
}

Antes de realizar las modificaciones la respuesta de la consulta a gemini es de este tipo:

```sql
-- DML statements [Dummy data]
INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Teléfono móvil', 'Smartphone de última generación', 799, 12, 'modelo1', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Camiseta', 'Camiseta de algodón', 20, 24, 'modelo2', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Sartén antiadherente', 'Sartén para cocinar', 35, 1, 'modelo3', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Balón de fútbol', 'Balón oficial de la FIFA', 50, 6, 'modelo4', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Muñeca', 'Muñeca de peluche para niños', 15, 10, 'modelo5', DATETIME('now'), DATETIME('now'));
```

## Estructura JSON para los dos primeros inserts

```json
"nombresito": "Teléfono móvil",
"descripcionsita": "Smartphone de última generación",
"precioaquel": 799,
"cantidadparce": 12,
"randomoelo": "modelo1"

"nombresito": "Camiseta",
"descripcionsita": "Camiseta de algodón",
"precioaquel": 20,
"cantidadparce": 24,
"randomoelo": "modelo2"
```

En resumen es conservar la logica que extrae los inserts, ya que gemini ya entrega la respuesta de la data generada.
Se requiere extraer las porciones de JSON de la respuesta de gemini y que la funcio AddDummyData retorne los 3 strings generados por gemini (class string, classMetadata [][]string) DummyDataResult .

1. 5 inserts
2. Porcion de JSON 1
3. Porcion de JSON 2

Podrias darme la respuesta en español