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

func Generate(projectName, dbType, configFile string) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)

	class, classMetadata, err := readConfigMetadata(configFile)
	if err != nil {
		fmt.Printf("Error leyendo el archivo de configuración: %s\n", err)
		fmt.Println("la clase es:", class)
		os.Exit(1)
	}
	fmt.Printf("Configuración leída: %+v\n %+v\n", class, classMetadata)
	// Aquí puedes agregar la lógica para generar el proyecto

	// fmt.Println("En generate\n")
	// fmt.Println("Class metadata", classMetadata)
	// longitud := len(classMetadata)
	// fmt.Println("longitud del map es:", longitud)

	generateClassTags(class, classMetadata)

	modifyBaseTemplates()

	// Generate folder structure
	createFolderStructure(projectName, class, classMetadata)

}

// func leerConfig(configFile string) ([]models.Tipo, error) {
// Por ahora solo leera un objeto JSON entonces la funcion retornara un map en la informacion de una clase
func readConfigMetadata(configFile string) (string, map[string]string, error) {
	jsonData, err := os.Open(configFile)
	if err != nil {
		return "", nil, err
	}
	defer jsonData.Close()

	// fmt.Println("JSONDATA ES:", jsonData)

	bytes, err := ioutil.ReadAll(jsonData)
	if err != nil {
		return "", nil, err
	}

	var tipos []models.Tipo
	if err := json.Unmarshal(bytes, &tipos); err != nil {
		return "", nil, err
	}

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

// func createModels() {

// }
