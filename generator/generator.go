package generator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/diego-all/run-from-gh/models"
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
	//config, err := readConfig(configFile)
	class, config, err := leerConfig(configFile)
	if err != nil {
		fmt.Printf("Error leyendo el archivo de configuración: %s\n", err)
		fmt.Println("la clase es:", class)
		os.Exit(1)
	}
	fmt.Printf("Configuración leída: %+v\n", config)
	// Aquí puedes agregar la lógica para generar el proyecto

	fmt.Println("extraer data")

	// for _, entity := range config {
	// 	fmt.Printf("Tipo: %s\n", entity.Tipo)
	// 	for atributo, detalles := range entity.Atributos {
	// 		fmt.Printf(" Atributo: %s, Tipo de Dato: %s\n", atributo, detalles.TipoDato)
	// 	}

	// }

	fmt.Println("iterar map")
	for k, v := range config {
		fmt.Printf("key[%s] value[%s]\n", k, v)
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

// ESTE ES EL QUE ANDO UNIENDO CON EL DE G3NERATOR
// func leerConfig(configFile string) ([]Entity, error) {
// func leerConfig(configFile string) ([]models.Tipo, error) {
// Por ahora solo leera un objeto JSON entonces la funcion retornara un map en la informacion de una clase
func leerConfig(configFile string) (string, map[string]string, error) {
	jsonData, err := os.Open(configFile)
	if err != nil {
		return "", nil, err
	}
	defer jsonData.Close()

	fmt.Println("JSONDATA ES:", jsonData)

	bytes, err := ioutil.ReadAll(jsonData)
	if err != nil {
		return "", nil, err
	}

	var tipos []models.Tipo
	if err := json.Unmarshal(bytes, &tipos); err != nil {
		return "", nil, err
	}

	fmt.Println("TIPOS ES:", tipos)

	// PROVISIONAL [Solo 1 Tipo del JSON]
	mapAtributos := make(map[string]string)
	var Class string // Declaración de la variable Class

	// Iterar sobre cada tipo y sus atributos
	for _, tipo := range tipos {
		Class = tipo.Tipo
		fmt.Println("Clase:", tipo.Tipo)
		fmt.Println("Atributos:")
		for nombreAtributo, atributo := range tipo.Atributos {

			fmt.Printf(" - %s: %s\n", nombreAtributo, atributo.TipoDato)

			// PROVISIONAL [Solo 1 Tipo del JSON]
			mapAtributos[nombreAtributo] = atributo.TipoDato
		}

		// PROVISIONAL [Solo 1 Tipo del JSON]
		oneType := true
		if oneType == true {
			break
		}
	}

	// PROVISIONAL [Solo 1 Tipo del JSON]
	fmt.Println("mapAtributos es: ", mapAtributos)

	return Class, mapAtributos, nil
}

// func createFolderStructure() {

// }

// func createModels() {

// }
