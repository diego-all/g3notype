package extractor

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func CallPythonExtractor(jsonPath string) (bytes.Buffer, error) {

	// Run the Python script
	// validate if the combinedoutput can be used
	cmd := exec.Command("python3", "extractor/readMap.py", jsonPath)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		//return out, fmt.Errorf("error ejecutando el script de Python: %v", err)
		return out, fmt.Errorf("error ejecutando el script de Python: %v, salida de error: %s", err, stderr.String())
	}

	return out, err
}

func ParseData(buffer bytes.Buffer) (string, [][]string, error) {

	// Read python script output
	output := buffer.String()
	lines := strings.Split(output, "\n")
	//fmt.Println("Lines: \n", lines, "\n")

	// The first element is the string type
	if len(lines) < 1 {
		return "", nil, fmt.Errorf("no hay suficiente salida del script de Python")
	}
	tipo := lines[0]

	// The following elements are the list of lists
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
