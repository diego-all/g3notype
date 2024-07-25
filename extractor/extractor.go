package extractor

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// func CallPythonExtractor(jsonPath string) (string, [][]string, error) {

// 	var aux string
// 	var matriz [][]string

// 	// pendiente capturar el error
// 	aux, matriz = exec.Command("python3", "extractor/readMap.py", jsonPath)

// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		fmt.Printf("Error: %s\n", err)
// 		fmt.Printf("Salida del comando: %s\n", string(output))
// 	}

// 	return string(output), matriz, err
// }

func CallPythonExtractor(jsonPath string) (string, [][]string, error) {
	// Ejecutar el script de Python
	// validar si se peude utilizar la combinedoutput
	cmd := exec.Command("python3", "extractor/readMap.py", jsonPath)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", nil, fmt.Errorf("error ejecutando el script de Python: %v", err)
	}

	// Leer la salida
	output := out.String()
	lines := strings.Split(output, "\n")

	// El primer elemento es el string tipo
	if len(lines) < 1 {
		return "", nil, fmt.Errorf("no hay suficiente salida del script de Python")
	}
	tipo := lines[0]

	// Los elementos siguientes son la lista de listas
	var matrizAtributos [][]string
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}
		atributo := strings.Split(line, "|")
		if len(atributo) == 2 {
			matrizAtributos = append(matrizAtributos, atributo)
		}
	}

	return tipo, matrizAtributos, nil
}
