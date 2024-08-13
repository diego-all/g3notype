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

		// Imprimir los resultados
		fmt.Printf("Tipo: %s\n", tipo)
		fmt.Println()
		for _, atributo := range matrizAtributos {
			fmt.Printf("Atributo: %s, Tipo de dato: %s\n", atributo[0], atributo[1])

		}

		configuration := models.Config{
			ProjectName: projectName,
			//Database: db
			Tipo:            tipo,
			MatrizAtributos: matrizAtributos,
		}

		// Improvement: Send the entire configuration directly from the struct
		generator.Generate(projectName, db, configuration, dummy)

		elapsed := time.Since(start)

		fmt.Printf("El tiempo de ejecución es: %s\n", elapsed)
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
