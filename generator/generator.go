package generator

import (
	"fmt"

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

func Generate(projectName string, dbType string, config models.Config, dummy bool) {
	fmt.Printf("Generando proyecto '%s' con base de datos '%s'\n", projectName, dbType)

	fmt.Println("CONFIG from Generatex (output python): ", config)

	for _, trin := range config.MatrizAtributos {

		fmt.Println(trin)
	}

	fmt.Println(config.Tipo)

	tiposGenerados := generateClassTags(config.Tipo, config.MatrizAtributos)
	fmt.Println("Longitud de tiposGenerados: (generator/Generate)", len(tiposGenerados))
	//fmt.Println("TIPO GENERADO:", tipoGenerado) // el mismo del retorno de la funcion
	fmt.Println("\n")

	// SUGERENCIA: OBTENER VALOR POR VALOR Y LLENAR  data := TemplateData{} para sustituir las plantillas, quizas se requieran archivos intermedios.
	generatedDatabaseDDL := generateDatabaseDDL(config.Tipo, config.MatrizAtributos, dummy)
	fmt.Println("El DDL es: \n", generatedDatabaseDDL)

	generatedModels := generateEntityModels(config.Tipo, config.MatrizAtributos)
	fmt.Println("Generated Models es: ", generatedModels)
	//TRIN TRINfmt.Println("Tipo del mapa:", reflect.TypeOf(generatedModels))
	//fmt.Println("Se han generado ", generatedModels, "EntityModels")
	fmt.Println("\n")

	modifyBaseTemplates(tiposGenerados) // Pueden variar
	//modifyBaseTemplates(generatedModels) // Pueden variar

	//SE TUESTA MIRAR SI UN SLEEP O VALIDAR BIEN

	// Generate folder structure
	//createFolderStructure(projectName, class, classMetadata, generateClassTags(class, classMetadata)) //recordar que no funciono mandando una funcion pero si el valor , tipoGenerado
	createFolderStructure(projectName, config.Tipo, config.MatrizAtributos)

}
