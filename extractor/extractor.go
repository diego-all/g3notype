package extractor

import (
	"fmt"
	"os/exec"
)

func CallPythonExtractor(jsonPath string) (string, error) {
	cmd := exec.Command("python3", "extractor/readMap.py", jsonPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		fmt.Printf("Salida del comando: %s\n", string(output))
	}
	return string(output), err
}
