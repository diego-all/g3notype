package models

// Estructura para los atributos
type Atributo struct {
	TipoDato string `json:"tipoDato"`
}

// Estructura para el tipo
type Tipo struct {
	Tipo      string              `json:"tipo"`
	Atributos map[string]Atributo `json:"atributos"`
}

type Config struct {
	ProjectName string
	//Database db
	Tipo            string
	MatrizAtributos [][]string
}

type DummyDataResult struct {
	Inserts    string
	CreateJSON string
	UpdateJSON string
}
