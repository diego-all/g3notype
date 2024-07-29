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

// init use a posicional argument (projectName)

// Pendiente intentar pasar data structure to data structure (list of list to string' slice )

var initCmd = &cobra.Command{
	Use:   "init [nombre del proyecto]",
	Short: "Inicializa un nuevo proyecto",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		start := time.Now()

		projectName := args[0]
		// generator.Generate(projectName, db)

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

		//Minimizar enviar todo directamente al struct

		generator.Generatex(projectName, db, configuration)
		//generator.Generate(projectName, db, jsonPath)

		elapsed := time.Since(start)

		fmt.Printf("El tiempo de ejecución es: %s\n", elapsed)
	},
}

// func init() {
// 	rootCmd.AddCommand(initCmd)
// 	initCmd.Flags().StringVarP(&db, "db", "d", "", "Tipo de base de datos (requerido)")
// 	//initCmd.Flags().StringVarP(&jsonPath, "config", "c", "", "Ruta del archivo JSON de configuración (opcional)")
// 	initCmd.Flags().StringVarP(&jsonPath, "config", "c", "inputs/classes.json", "Ruta del archivo JSON de configuración (opcional)")
// 	initCmd.MarkFlagRequired("db")
// }

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&db, "db", "d", "", "Tipo de base de datos (requerido)")
	initCmd.MarkFlagRequired("db")
	// Validar "" o "/inputs/classes.json"
	//initCmd.Flags().StringVarP(&jsonPath, "config", "c", "inputs/classes.json", "Ruta del archivo JSON de configuración")
	initCmd.Flags().StringVarP(&jsonPath, "config", "c", "", "Ruta del archivo JSON de configuración")
}
