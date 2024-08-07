tengo una funcion en golang que recibe como parametros un string y una matriz y retorna un string:


func AddDummyData(class string, classMetadata [][]string) string {

	dummyData := GenerateDummyData(class, classMetadata)

	return ExtractInsertStatements(dummyData)
}


En GenerateDummyData() hay una logica que hace una consulta a traves de la API de gemini,

func GenerateDummyData(class string, classMetadata [][]string) string {
//Logica para hacer consulta a gemini
	return fmt.Sprintf("%s", strings.Join(parts, "\n"))
}

En dicha consulta (query) requiere contatenar dos variables:

class que es un string y classMetadata que es una matriz de strings.


Actualmente el script lo que hace es construir 5 insert para una base de datos sqlite.

Luego de esto hay una logica que a partir de una funcion con una expresion regular extrae los 5 insert y los retorna en forma de string.


Ahora requiero modificar esta consulta para que gemini genere dummy data en forma o structura de JSON considerando los tipos de dato incluidos en los metadatos que recibe la funcion.

Podrias ayudarme a modificar esta logica para que ahora con esta segunda parte que genera dummydata para construir el body de un JSON (considerando que siguen siendo strings) puedan ser extraidas. Una puede ser a partir del primer insert emularia una peticion request de CREATE y a partir del segundo insert crear un UPDATE.

AHora luego de extraer los 3 valores solicitados:

1. 5 inserts en un string
2. 1 string que emule la porcion de un JSON create a partir del primer insert.
3. 1 string que emule la porcion de un JSON update a partir del segundo insert

Quiza AHora las funciones deban ser modificadas y su valor de retorno no seria solo un string como esta actualmente, sino otra estructura como un array o un map.

Considerando nuevamente no retornar 3 strings sino un struct con los 3 strings.


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

Podrias darme la respuesta en español


