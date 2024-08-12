Tengo un script en golang que realiza una consulta a gemini y obtiene el siguiente response:

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

## Estructura JSON para los 2 primeros inserts:

**createBody:**

```
"nombresito": "Telefono movil",
"descripcionsita": "Smartphone de ultima generacion",
"precioaquel": 799,
"cantidadparce": 12,
"randomoelo": "ModeloA"
```

**updateBody:**

```
"nombresito": "Camiseta",
"descripcionsita": "Camiseta de algodon",
"precioaquel": 20,
"cantidadparce": 6,
"randomoelo": "ModeloB"
``` 


Esta es la estructura del script

package generator

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


func ExtractInsertStatements(data string) string {
	re := regexp.MustCompile(`(?i)INSERT INTO [^\;]+;`)
	inserts := re.FindAllString(data, -1)
	return strings.Join(inserts, "\n")
}

func AddDummyData(class string, classMetadata [][]string) DummyDataResult {
	dummyData := GenerateDummyData(class, classMetadata)
	createJSON, updateJSON := ExtractUpsertCollection(dummyData)

	return models.DummyDataResult{
		Inserts:    dummyData,
		CreateJSON: "en construccion",
		UpdateJSON: "en construccion",
	}
}


Necesito que me ayudes a generar la logica para extraer los dos strings sin utilizar ExtractUpsertCollection() enviando dummydata como parametro ya que solo trae la data de los inserts.
Requiero que sea extraida de la response