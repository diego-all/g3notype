Estoy intentando ejecutar un script en golang:

root@pho3nix:/home/diegoall/MAESTRIA_ING/CLI/TEST-G3N# go run github.com/diego-all/run-from-gh@latest init --db sqlite --dummy --config product.json exampleAPI
Error al ejecutar el script de Python: error ejecutando el script de Python: exit status 2

Internamente el script tiene una funcionalidad en python basica.

Podrias ayudarme a corregir el error por favor? Es una aplicacion de tipo CLI.

/cmd/init.go

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
	dummy    bool
)

// init use a posicional argument (projectName)

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

		// Parse the data
		tipo, matrizAtributos, err := extractor.ParseData(buffer)
		if err != nil {
			fmt.Printf("Error al parsear los datos: %v\n", err)
			return
		}

		// //Imprimir los resultados
		// fmt.Printf("Tipo: %s\n", tipo)
		// fmt.Println()
		// for _, atributo := range matrizAtributos {
		// 	fmt.Printf("Atributo: %s, Tipo de dato: %s\n", atributo[0], atributo[1])
		// }
		fmt.Print("\n")
		fmt.Println("Leyendo configuración ... \n")

		configuration := models.Config{
			ProjectName: projectName,
			//Database: db
			Tipo:            tipo,
			MatrizAtributos: matrizAtributos,
		}

		// Improvement: Send the entire configuration directly from the struct
		generator.Generate(projectName, db, configuration, dummy)

		elapsed := time.Since(start)

		//fmt.Print("\n")
		fmt.Printf("El tiempo de ejecución es: %s\n", elapsed)
		fmt.Print("\n")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&db, "db", "d", "", "Tipo de base de datos (requerido)")
	initCmd.MarkFlagRequired("db")
	//initCmd.Flags().StringVarP(&jsonPath, "config", "c", "inputs/classes.json", "Ruta del archivo JSON de configuración")
	initCmd.Flags().StringVarP(&jsonPath, "config", "c", "", "Ruta del archivo JSON de configuración")
	initCmd.Flags().BoolVarP(&dummy, "dummy", "u", false, "Generar Dummy data usando Gemini (Requiere API Key)")
}

/cmd/root.go

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "g3notype",
	Short: "CLI application to generate secure REST API projects from a domain model",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	// Initialize flags and global settings if necessary

}

/extractor/extractor.go

package extractor

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func CallPythonExtractor(jsonPath string) (bytes.Buffer, error) {

	// Run the Python script
	// validate if the combinedoutput can be used
	cmd := exec.Command("python3", "extractor/readMap.py", jsonPath)

	var out bytes.Buffer

	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return out, fmt.Errorf("error ejecutando el script de Python: %v", err)
	}

	return out, err
}

func ParseData(buffer bytes.Buffer) (string, [][]string, error) {

	// Read python script output
	output := buffer.String()
	lines := strings.Split(output, "\n")
	//fmt.Println("Lines: \n", lines, "\n")

	// The first element is the string type
	if len(lines) < 1 {
		return "", nil, fmt.Errorf("no hay suficiente salida del script de Python")
	}
	tipo := lines[0]

	// The following elements are the list of lists
	var matrizAtributos [][]string
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		atributo := strings.Split(line, "|")
		if len(atributo) == 2 {
			matrizAtributos = append(matrizAtributos, atributo)
		}
	}

	return tipo, matrizAtributos, nil

}

/extractor/readMap.py

import sys
import json

def main(json_path):
    with open(json_path, 'r') as archivo:
        datos_lista = json.load(archivo)
        datos = datos_lista[0]  # Primer elemento del array JSON
        tipo = datos['tipo']
        atributos = datos['atributos']

        matriz_atributos = []
        for nombre, detalle in atributos.items():
            tipo_dato = detalle['tipoDato']
            matriz_atributos.append([nombre, tipo_dato])

        return tipo, matriz_atributos

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Uso: python3 extractor/readMap.py <ruta_del_json>")
        sys.exit(1)
    json_path = sys.argv[1]
    tipo, matriz_atributos = main(json_path)
    
    # Salida como string y lista de listas
    print(tipo)
    for atributo in matriz_atributos:
        print(f"{atributo[0]}|{atributo[1]}")

/generator/generator.go

package generator

import (
	"fmt"

	"github.com/diego-all/run-from-gh/models"
)

type Attribute struct {
	TipoDato string json:"tipoDato"
}

type Entity struct {
	Tipo      string               json:"tipo"
	Atributos map[string]Attribute json:"atributos"
}

func Generate(projectName string, dbType string, config models.Config, dummy bool) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)
	fmt.Print("\n")

	// fmt.Println("Config from Generate (output python): ", config)

	// for _, trin := range config.MatrizAtributos {
	// 	fmt.Println(trin)
	// }

	//fmt.Println(config.Tipo)

	// The NaturalID used in generateEntityModels is calculated
	tiposGenerados := generateHandlers(config.Tipo, config.MatrizAtributos)
	//fmt.Println("Longitud de tiposGenerados: (generator/Generate)", len(tiposGenerados))

	//generatedDatabaseDDL := generateDatabaseDDL(config.Tipo, config.MatrizAtributos, dummy)
	//fmt.Println("El DDL es: \n", generatedDatabaseDDL)
	generateDatabaseDDL(config.Tipo, config.MatrizAtributos, dummy)

	//generatedModels := generateEntityModels(config.Tipo, config.MatrizAtributos)
	generateEntityModels(config.Tipo, config.MatrizAtributos)
	//fmt.Println("Generated Models es: ", generatedModels)

	modifyBaseTemplates(tiposGenerados)

	// Generate folder structure
	createFolderStructure(projectName, config.Tipo, config.MatrizAtributos)

}

/models/models.go

package models

// Estructura para los atributos
type Atributo struct {
	TipoDato string json:"tipoDato"
}

// Estructura para el tipo
type Tipo struct {
	Tipo      string              json:"tipo"
	Atributos map[string]Atributo json:"atributos"
}

type Config struct {
	ProjectName string
	//Database db
	Tipo            string
	MatrizAtributos [][]string
}

type DummyDataResult struct {
	Inserts    string
	CreateJSON string
	UpdateJSON string
}

/inputs/product.json

[
    {
      "tipo": "Product",
      "atributos": {
        "name": {
          "tipoDato": "string"
        },
        "description": {
          "tipoDato": "string"
        },
        "price": {
          "tipoDato": "int"
        },
        "quantity": {
          "tipoDato": "int"
        }
      }
    }
]

go.mod

module github.com/diego-all/run-from-gh

go 1.21

//toolchain go1.22.1

// go 1.22.1

require (
	github.com/google/generative-ai-go v0.17.0
	github.com/joho/godotenv v1.5.1
	github.com/spf13/cobra v1.8.1
	google.golang.org/api v0.189.0
)

require (
	cloud.google.com/go v0.115.0 // indirect
	cloud.google.com/go/ai v0.8.2 // indirect
	cloud.google.com/go/auth v0.7.2 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.3 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	cloud.google.com/go/longrunning v0.5.11 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.13.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.53.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0 // indirect
	go.opentelemetry.io/otel v1.28.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	golang.org/x/crypto v0.25.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/oauth2 v0.21.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240725223205-93522f1f2a9f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240725223205-93522f1f2a9f // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

Podrias darme la respuesta en español
