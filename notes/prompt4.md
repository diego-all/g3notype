Tengo el siguiente script en golang y requiero crear una estructura que sea visible desde otro paquete.

Especificamente el struct Config,  que sea parametro de entrada de la funcion generator.Generatex(projectName, db, tipo, Config)


Podrias ayudarme a construirlo ya que al asignarlo como parametro me aparece:

Config (type) is not an expressioncompiler

Aca esta init.go

package cmd

import (
	"fmt"

	"github.com/diego-all/run-from-gh/extractor"
	"github.com/diego-all/run-from-gh/generator"
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

		type Config struct {
			tipo            string
			matrizAtributos [][]string
		}

		Configuration = Config{
			tipo:            tipo,
			matrizAtributos: matrizAtributos,
		}

		generator.Generatex(projectName, db, tipo, Configuration)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&db, "db", "d", "", "Tipo de base de datos (requerido)")
	initCmd.MarkFlagRequired("db")
	initCmd.Flags().StringVarP(&jsonPath, "config", "c", "", "Ruta del archivo JSON de configuración")
}

requiero poder invocar la funcion generator.Generatex(projectName, db, tipo, Configuration) Pero hay un problema con el parametro Configuration
Por otra parte en el paquete generator en la funcion Generatex el ultimo parametro Configuration tambien me arroja un error:

mixed named and unnamed parameterssyntax
var Config invalid type


package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/diego-all/run-from-gh/models"
)

// EXTRAE EN MAP
type Attribute struct {
	TipoDato string `json:"tipoDato"`
}

type Entity struct {
	Tipo      string               `json:"tipo"`
	Atributos map[string]Attribute `json:"atributos"`
}


func Generatex(projectName, dbType, configFile string, Config) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)

	fmt.Println("CONFIGFILE from Generatex (output python): ", configFile)

	class, classMetadata, err := readConfigMetadatax(configFile)
	if err != nil {
		fmt.Printf("Error leyendo el archivo de configuración: %s\n", err)
		fmt.Println("la clase es:", class)
		os.Exit(1)
	}
	fmt.Printf("Configuración leída: %+v\n %+v\n", class, classMetadata)

	fmt.Println("\n")

	fmt.Println("\n")


}

AL parecer debo emplear alguna estrategia ṕara que el type COnfig sea visible desde el paquete generator.


Podrias ayudarme con esto por favor y darme la respuesta en español.