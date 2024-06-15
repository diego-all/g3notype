package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "my-cli-app",
	Short: "Una aplicación CLI para generar proyectos",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Aquí puedes inicializar flags y configuraciones globales si es necesario
}
