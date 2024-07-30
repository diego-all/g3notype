Tengo el siguiente App CLI en golang, Requiero hacer una modificacion para que el comando de ejecucion al recibir otro parametro ejecute otra funcionalidad:

go run main.go init --db sqlite --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest

Dicho nuevo parametro es --dummy, la idea es que este nuevo parametro al estar presente en el comando de ejecucion permita ejecutar una funcionalidad extra. 

go run main.go init --db sqlite --dummy --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest

El programa esta estructurado de la siguiente forma:

init.go

package cmd

import (
	"fmt"
	"time"

	"github.com/diego-all/run-from-gh/extractor"
	"github.com/diego-all/run-from-gh/generator"
	"github.com/diego-all/run-from-gh/models"
	"github.com/spf13/cobra"
)

var (
	db       string
	jsonPath string
)

var initCmd = &cobra.Command{
	Use:   "init [nombre del proyecto]",
	Short: "Inicializa un nuevo proyecto",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()

		projectName := args[0]

		buffer, err := extractor.CallPythonExtractor(jsonPath)
		if err != nil {
			fmt.Printf("Error al ejecutar el script de Python: %v\n", err)
			return
		}

		// Parsear los datos
		tipo, matrizAtributos, err := extractor.ParseData(buffer)
		if err != nil {
			fmt.Printf("Error al parsear los datos: %v\n", err)
			return
		}

		// Imprimir los resultados
		fmt.Printf("Tipo: %s\n", tipo)
		fmt.Println()
		for _, atributo := range matrizAtributos {
			fmt.Printf("Atributo: %s, Tipo de dato: %s\n", atributo[0], atributo[1])

		}

		configuration := models.Config{
			//ProjectName: projectName,
			//Database: db
			Tipo:            tipo,
			MatrizAtributos: matrizAtributos,
		}

		generator.Generatex(projectName, db, configuration)

		generator.GenerateDummyData(configuration)

		elapsed := time.Since(start)

		fmt.Printf("El tiempo de ejecución es: %s\n", elapsed)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&db, "db", "d", "", "Tipo de base de datos (requerido)")
	initCmd.MarkFlagRequired("db")
	initCmd.Flags().StringVarP(&jsonPath, "config", "c", "", "Ruta del archivo JSON de configuración")
}

root.go

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "Una aplicación CLI para generar proyectos",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Aquí puedes inicializar flags y configuraciones globales si es necesario
}

generator.go

package generator

import (
	"fmt"

	"github.com/diego-all/run-from-gh/models"
)

type Attribute struct {
	TipoDato string `json:"tipoDato"`
}

type Entity struct {
	Tipo      string               `json:"tipo"`
	Atributos map[string]Attribute `json:"atributos"`
}


func Generatex(projectName string, dbType string, config models.Config) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)

	fmt.Println("CONFIG from Generatex (output python): ", config)

	for _, trin := range config.MatrizAtributos {

		fmt.Println(trin)
	}

	fmt.Println(config.Tipo)

	tiposGenerados := generateClassTags(config.Tipo, config.MatrizAtributos)
	fmt.Println("Longitud de tiposGenerados: (generator/Generate)", len(tiposGenerados))
	fmt.Println("\n")

	generatedDDL := generateDDLStatement(config.Tipo, config.MatrizAtributos)
	fmt.Println("El DDL es: \n", generatedDDL)

	generatedModels := generateEntityModels(config.Tipo, config.MatrizAtributos)
	fmt.Println("Generated Models es: ", generatedModels)
	fmt.Println("\n")

	modifyBaseTemplates(tiposGenerados)

	createFolderStructure(projectName, config.Tipo, config.MatrizAtributos)

}

dummyData.go

package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

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

func GenerateDummyData(config models.Config) {
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

	// Definir la consulta
	query := `Tengo un modelo de datos: Books con los siguientes atributos y su tipo de dato correspondiente
		nombre | string
		descripcion| string
		precio | int
		cantidad | int
		random| int
		created_at|DATETIME('now')
		updated_at|DATETIME('now')
		
		Requiero construir basado en los datos anteriores las sentencias insert con data dummy, en total 5 sentencias para una base de datos sqlite, como las siguientes:
		
		-- DML statements [Dummy data]
		-- INSERT INTO products (name, description, price, created_at, updated_at)
		--     VALUES ('Teléfono móvil', 'Smartphone de última generación', 799, DATETIME('now'), DATETIME('now'));
		
		-- INSERT INTO products (name, description, price, created_at, updated_at)
		--     VALUES ('Camiseta', 'Camiseta de algodón', 20, DATETIME('now'), DATETIME('now'));
		
		-- INSERT INTO products (name, description, price, created_at, updated_at)
		--     VALUES ('Sartén antiadherente', 'Sartén para cocinar', 35, DATETIME('now'), DATETIME('now'));
		
		-- INSERT INTO products (name, description, price, created_at, updated_at)
		--     VALUES ('Balón de fútbol', 'Balón oficial de la FIFA', 50, DATETIME('now'), DATETIME('now'));
		
		-- INSERT INTO products (name, description, price, created_at, updated_at)
		--     VALUES ('Muñeca', 'Muñeca de peluche para niños', 15, DATETIME('now'), DATETIME('now'));`

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

	// Imprimir el contenido de Parts
	for _, candidate := range data.Candidates {
		for _, part := range candidate.Content.Parts {
			fmt.Println(part)
		}
	}
}

main.go

package main

import "github.com/diego-all/run-from-gh/cmd"

func main() {
	//cmd.Execute()
	cmd.Execute()
}

La idea es que al suministrar en el comando --dummy ejecutar la funcion GenerateDummyData(), de lo contrario no ejecutarla.

Podrias darme la respuesta en español por favor.