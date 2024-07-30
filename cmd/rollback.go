package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Restaura los archivos genéricos a partir de los archivos base",
	Run: func(cmd *cobra.Command, args []string) {
		rollbackFiles := []struct {
			src, dst string
		}{
			{"base-templates/cmd/api/handlers-entity-base.txt", "base-templates/cmd/api/handlers-entity-generic.txt"},
			{"base-templates/database/up.sql-base.txt", "base-templates/database/up.sql-generic.txt"},
			{"base-templates/internal/entities-base.txt", "base-templates/internal/entities-generic.txt"},
		}

		for _, file := range rollbackFiles {
			err := copyFile(file.src, file.dst)
			if err != nil {
				fmt.Printf("Error al copiar %s a %s: %v\n", file.src, file.dst, err)
				return
			}
		}

		allFilesMatch := true
		for _, file := range rollbackFiles {
			match, err := compareFiles(file.src, file.dst)
			if err != nil {
				fmt.Printf("Error al comparar %s y %s: %v\n", file.src, file.dst, err)
				return
			}
			if !match {
				allFilesMatch = false
				fmt.Printf("El contenido de %s y %s no coincide\n", file.src, file.dst)
			}
		}

		if allFilesMatch {
			fmt.Println("El contenido de las pretemplates ha sido restablecido correctamente.")
		} else {
			fmt.Println("Algunos archivos no se copiaron correctamente.")
		}
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}

func copyFile(src, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func compareFiles(file1, file2 string) (bool, error) {
	content1, err := ioutil.ReadFile(file1)
	if err != nil {
		return false, err
	}
	content2, err := ioutil.ReadFile(file2)
	if err != nil {
		return false, err
	}
	return string(content1) == string(content2), nil
}
