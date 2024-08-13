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

func Generate(projectName string, dbType string, config models.Config, dummy bool) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)

	fmt.Println("CONFIG from Generatex (output python): ", config)

	for _, trin := range config.MatrizAtributos {

		fmt.Println(trin)
	}

	fmt.Println(config.Tipo)

	tiposGenerados := generateClassTags(config.Tipo, config.MatrizAtributos)
	fmt.Println("Longitud de tiposGenerados: (generator/Generate)", len(tiposGenerados))

	fmt.Println("\n")

	generatedDatabaseDDL := generateDatabaseDDL(config.Tipo, config.MatrizAtributos, dummy)
	fmt.Println("El DDL es: \n", generatedDatabaseDDL)

	generatedModels := generateEntityModels(config.Tipo, config.MatrizAtributos)
	fmt.Println("Generated Models es: ", generatedModels)

	modifyBaseTemplates(tiposGenerados)

	// Generate folder structure
	createFolderStructure(projectName, config.Tipo, config.MatrizAtributos)

}
