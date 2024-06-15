package cmd

import (
	"github.com/diego-all/run-from-gh/generator"
	"github.com/spf13/cobra"
)

var db string

var initCmd = &cobra.Command{
	Use:   "init [nombre del proyecto]",
	Short: "Inicializa un nuevo proyecto",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		// generator.Generate(projectName, db)
		generator.Generate(projectName, db)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&db, "db", "d", "", "Tipo de base de datos (requerido)")
	initCmd.MarkFlagRequired("db")
}
