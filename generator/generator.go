package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// func Generate(projectName, dbType string, jsonPath string) {
// 	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)

// 	if jsonPath != "" {
// 		// Leer el archivo JSON
// 		file, err := os.Open(jsonPath)
// 		fmt.Println("FILE ES:", file)
// 		if err != nil {
// 			fmt.Printf("Error al abrir el archivo JSON: %v\n", err)
// 			return
// 		}
// 		defer file.Close()

// 		byteValue, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			fmt.Printf("Error al leer el archivo JSON: %v\n", err)
// 			return
// 		}

// 		var config map[string]interface{}
// 		if err := json.Unmarshal(byteValue, &config); err != nil {
// 			fmt.Printf("Error al parsear el archivo JSON: %v\n", err)
// 			return
// 		}

// 		fmt.Printf("Configuración leída del archivo JSON: %v\n", config)
// 		// Aquí puedes agregar la lógica para utilizar la configuración leída
// 	}

// 	// Aquí puedes agregar la lógica para generar el proyecto
// }

// GENERATE() ANTIGUO SIN --CONFIG
// func Generate(projectName, dbType, configFile string) {
// 	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)
// 	config, err := readConfig(configFile)
// 	if err != nil {
// 		fmt.Printf("Error leyendo el archivo de configuración: %s\n", err)
// 		os.Exit(1)
// 	}
// 	fmt.Printf("Configuración leída: %+v\n", config)
// 	// Aquí puedes agregar la lógica para generar el proyecto
// }

// READCONFIG() ANTIGUO SIN --CONFIG
// func readConfig(configFile string) (map[string]interface{}, error) {
// 	file, err := os.Open(configFile)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	bytes, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var config map[string]interface{}
// 	if err := json.Unmarshal(bytes, &config); err != nil {
// 		return nil, err
// 	}

// 	return config, nil
// }

// EXTRAE EN MAP
type Attribute struct {
	TipoDato string `json:"tipoDato"`
}

type Entity struct {
	Tipo      string               `json:"tipo"`
	Atributos map[string]Attribute `json:"atributos"`
}

func Generate(projectName, dbType, configFile string) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)

	fmt.Println("LEYENDO CONFIG")
	config, err := readConfig(configFile)
	if err != nil {
		fmt.Printf("Error leyendo el archivo de configuración: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Configuración leída: %+v\n", config)
	// Aquí puedes agregar la lógica para generar el proyecto

	fmt.Println("extraer data")

	for _, entity := range config {
		fmt.Printf("Tipo: %s\n", entity.Tipo)
		for atributo, detalles := range entity.Atributos {
			fmt.Printf(" Atributo: %s, Tipo de Dato: %s\n", atributo, detalles.TipoDato)
		}

	}

}

func readConfig(configFile string) ([]Entity, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var config []Entity
	if err := json.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return config, nil
}

// EXTRAE EN MAP
