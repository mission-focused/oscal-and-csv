/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/mission-focused/oscal-and-csv/src/pkg/common"
	"github.com/mission-focused/oscal-and-csv/src/pkg/oscal"
	"github.com/spf13/cobra"
)

// GenerateCmd represents the Generate command
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generates a CSV template for a specified OSCAL model",
	Long: `generates a CSV template for a specified OSCAL model
	Examples:
		oac generate [MODEL]`,
}

type CatalogFlags struct {
	OutputFile string // -o --output-file

}

var catalogOpts = &CatalogFlags{}

var CatalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Generate a catalog csv file",
	Run: func(cmd *cobra.Command, args []string) {
		var outputfile string

		if catalogOpts.OutputFile != "" {
			outputfile = catalogOpts.OutputFile
		} else {
			outputfile = "catalog.csv"
		}

		// Check for existing file - error if exists?
		if common.CheckFileExists(outputfile) {
			slog.Error(fmt.Sprintf("Error: %s already exists", outputfile))
			os.Exit(1)
		}

		template, err := oscal.CatalogToTemplate()
		if err != nil {
			slog.Error("Unable to generate a catalog template csv")
			os.Exit(1)
		}

		err = common.WriteCSV([][]string{template}, outputfile)
		if err != nil {
			slog.Error(fmt.Sprintf("Error: %v", err))
			os.Exit(1)
		}

		slog.Info(fmt.Sprintf("CSV template written to: %s", outputfile))

	},
}

func init() {
	rootCmd.AddCommand(GenerateCmd)
	GenerateCmd.AddCommand(CatalogCmd)

	// Here you will define your flags and configuration settings.
	CatalogCmd.Flags().StringVarP(&catalogOpts.OutputFile, "output-file", "o", "", "the path to write the output file")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GenerateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GenerateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
