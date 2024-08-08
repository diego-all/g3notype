Tengo esta funcion en golang: GenerateDummyData()

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

	// Unir las partes en una sola cadena de texto
	fmt.Println("GENERATEDUMMYDATA SENTENCIAS INSERT: \n")
	fmt.Println(fmt.Sprintf("%s", strings.Join(parts, "\n")))
	return fmt.Sprintf("%s", strings.Join(parts, "\n"))
}

Lo que imprime es lo siguiente:

GENERATEDUMMYDATA SENTENCIAS INSERT: 

## Sentencias INSERT con data dummy

```sql
-- DML statements [Dummy data]
INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Telefono movil', 'Smartphone de ultima generacion', 799, 24, 'modelo1', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Camiseta', 'Camiseta de algodon', 20, 12, 'modelo2', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Sarten antiadherente', 'Sarten para cocinar', 35, 8, 'modelo3', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Balon de futbol', 'Balon oficial de la FIFA', 50, 10, 'modelo4', DATETIME('now'), DATETIME('now'));

INSERT INTO products (nombresito, descripcionsita, precioaquel, cantidadparce, randomoelo, created_at, updated_at)
VALUES ('Muneca', 'Muneca de peluche para ninos', 15, 6, 'modelo5', DATETIME('now'), DATETIME('now'));
```

## Estructura JSON de los dos primeros inserts

```json
"nombresito": "Telefono movil",
"descripcionsita": "Smartphone de ultima generacion",
"precioaquel": 799,
"cantidadparce": 24,
"randomoelo": "modelo1"

"nombresito": "Camiseta",
"descripcionsita": "Camiseta de algodon",
"precioaquel": 20,
"cantidadparce": 12,
"randomoelo": "modelo2"
``` 

Requiero que me ayudes generando una logica que permita extraer los 5 inserts y asignarlos a un string, de igual forma con las otras 2 estructuras extraerlas y 
asignarlas a 2 strings independientes:

Por favor realiza la extraccion de la response de gemini.

Esta funcion es llamada desde esta funcion AddDummyData(), por favor crea las funciones necesarias para que desde AddDummyData puedan ser llamadas
y extraer la informacion que relaciono


func AddDummyData(class string, classMetadata [][]string) models.DummyDataResult {

	dummyData := GenerateDummyData(class, classMetadata)

	//llamar las funciones para obtener los 3 strings extraidos

	return DummyDataResult{
		Inserts:    "en construccion",
		CreateJSON: "en construccion",
		UpdateJSON: "en construccion",
	}
	
}

package models


type DummyDataResult struct {
	Inserts    string
	CreateJSON string
	UpdateJSON string
}


Te aclaro que por favor no modifiques la consulta hacia a gemini, para los 5 inserts se requiere conservar created_at y updated_at pero para los otros 
2 strings no. Reitero no modificar la query a gemini, gemini ya entrega la data generada solo es implementar la forma de extraerla y retronarla parseada.


Podrias darme la respuesta en español.