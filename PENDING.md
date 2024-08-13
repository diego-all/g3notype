# Pending

- Dummy data diferente a los inserts.
- Ajustar identaciones.
- Collections son en formato **json** curls <> Diferentes de postman e insomnia, estos requieren los verbos HTTP.
- Homologar tipos de datos en golang con respecto a la base de datos utilizada. (Data types)
- Generate dummy data a veces se equivoca con los modelos en el plural para crear la base de datos (ej. productos.json) o con el idioma español.
- Nombre del CLI (my-cli-app)
- Ajuste descriptivo (book.Name vs book.nombre) crear tipos (Nombre)
- Funciona sin repositorio. Analizar como hacer que funcione con repositorio y crearlo con git. (analizar como crea el repo git init y demas
    cmd/api/handlers-Book.go:8:2: no required module provides package github.com/diego-all/books-API/internal; to add it:
- createFolderStructure(projectName, class, classMetadata, generateClassTags(class, classMetadata)) //recordar que no funciono mandando una funcion pero si el valor , tipoGenerado
- -- DML statements [Dummy data] by Gemini
-- Tener cuidado con caracteres especiales que no acepta SQL lite como apostrofe por ejemplo para posesivo sajon
- Logica de request update sin gemini
- Al parecer es necesario crear una estructura temporal ya que TemplateData no puede modificarse en tiempo de ejecucion con el fin de generar los tipos para Request y Response
- Natural ID (Name)
- 	// //Error al ejecutar la plantilla: template: fileContent:8:2: executing "fileContent" at <.handlers_typeEntityRequest>: handlers_typeEntityRequest is an unexported field of struct type generator.preTemplateData
- 		//entity:  "{{.entity}}",   // NO funciona con minusculas seguir indagando
- 	// Pilas con "name,omitempty"`
- Debug mode with datatypes and generated data