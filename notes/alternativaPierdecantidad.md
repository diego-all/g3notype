

func AddDummyData(class string, classMetadata [][]string) models.DummyDataResult {
	// Llamar a GenerateDummyData y obtener la respuesta completa
	dummyData := GenerateDummyData(class, classMetadata)

	// Extraer las sentencias INSERT
	inserts := ExtractInserts(dummyData)

	// Extraer y validar los JSON bodies
	createJSON, updateJSON := ExtractJSONBodies(dummyData)
	createJSON, updateJSON = EnsureUniqueJSONBodies(createJSON, updateJSON)

	return models.DummyDataResult{
		Inserts:    inserts,
		CreateJSON: createJSON,
		UpdateJSON: updateJSON,
	}
}


func EnsureUniqueJSONBodies(createJSON, updateJSON string) (string, string) {
	for createJSON == updateJSON {
		// Repetimos la extracción hasta que sean diferentes.
		// Aquí puedes ajustar para generar nuevos datos.
		createJSON, updateJSON = ExtractJSONBodies(createJSON + updateJSON) // Ejemplo, podrías generar nuevos datos.
	}
	return createJSON, updateJSON
}


func ExtractJSONBodies(data string) (string, string) {
	re := regexp.MustCompile(`"(.+?)": "(.+?)",?`)
	matches := re.FindAllStringSubmatch(data, -1)

	if len(matches) == 0 {
		return "", ""
	}

	var createJSON, updateJSON strings.Builder
	for i, match := range matches {
		if i%2 == 0 {
			createJSON.WriteString(match[0] + "\n")
		} else {
			updateJSON.WriteString(match[0] + "\n")
		}
	}

	return createJSON.String(), updateJSON.String()
}


func ExtractInserts(data string) string {
	start := strings.Index(data, "INSERT INTO")
	if start == -1 {
		return ""
	}
	return data[start:]
}





