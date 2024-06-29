package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Generate(projectName, dbType string, jsonPath string) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)

	if jsonPath != "" {
		// Leer el archivo JSON
		file, err := os.Open(jsonPath)
		if err != nil {
			fmt.Printf("Error al abrir el archivo JSON: %v\n", err)
			return
		}
		defer file.Close()

		byteValue, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Printf("Error al leer el archivo JSON: %v\n", err)
			return
		}

		var config map[string]interface{}
		if err := json.Unmarshal(byteValue, &config); err != nil {
			fmt.Printf("Error al parsear el archivo JSON: %v\n", err)
			return
		}

		fmt.Printf("Configuración leída del archivo JSON: %v\n", config)
		// Aquí puedes agregar la lógica para utilizar la configuración leída
	}

	// Aquí puedes agregar la lógica para generar el proyecto
}
