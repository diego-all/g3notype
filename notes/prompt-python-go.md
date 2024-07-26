Requiero ejecutar un script de python desde una CLI de golang para extraer los valores de un JSON y conservar el orden ya que golang no permite esto de forma nativa.

Me esta apareciendo este error al ejecutar el programa. Podrias ayudarme a solucionarlo y darme la respuesta en español?

root@pho3nix:/home/diegoall/MAESTRIA_ING/CLI/run-from-gh# go run main.go init --db sqlite --config /home/diegoall/MAESTRIA_ING/CLI/run-from-gh/inputs/classes.json projectTest
Error al ejecutar el script de Python: exit status 2

package main

import "github.com/diego-all/run-from-gh/cmd"

func main() {
	//cmd.Execute()
	cmd.Execute()
}

package cmd

import (
	"github.com/diego-all/run-from-gh/extractor"
	"github.com/diego-all/run-from-gh/generator"
	"github.com/spf13/cobra"
)

var (
	db       string
	jsonPath string
)

// init use a posicional argument (projectName)

var initCmd = &cobra.Command{
	Use:   "init [nombre del proyecto]",
	Short: "Inicializa un nuevo proyecto",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		// generator.Generate(projectName, db)

		extractor.CallPythonExtractor()
		generator.Generate(projectName, db, jsonPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&db, "db", "d", "", "Tipo de base de datos (requerido)")
	initCmd.MarkFlagRequired("db")
	initCmd.Flags().StringVarP(&jsonPath, "config", "c", "", "Ruta del archivo JSON de configuración")
}

El script de python simplemente extrae data de un JSON:

import json

# Asumiendo que el JSON está en un archivo llamado 'classes.json' en la carpeta 'inputs'
with open('inputs/classes.json', 'r') as archivo:
    # Carga el JSON como una lista
    datos_lista = json.load(archivo)

    # Accede al primer objeto de la lista (si solo hay uno)
    datos = datos_lista[0]

    # Extraer el tipo
    tipo = datos['tipo']
    print(f'Tipo: {tipo}')

    # Extraer los atributos y sus tipos de dato
    atributos = datos['atributos']
    for nombre, detalle in atributos.items():
        tipo_dato = detalle['tipoDato']
        print(f'Atributo: {nombre}, Tipo de dato: {tipo_dato}')

El objetivo es lograr leer esta data antes de hacer el llamado a generator.Generate(projectName, db, jsonPath)

Y pasarle la data extraida del script de python a esta funcion en vez del jsonPath