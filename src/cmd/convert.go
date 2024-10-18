/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/mission-focused/oscal-and-csv/src/pkg/common"
	"github.com/mission-focused/oscal-and-csv/src/pkg/oscal"
	"github.com/spf13/cobra"
)

type flags struct {
	OutputFile string // -o --output-file
	InputFile  string // -f --input-file
}

var opts = &flags{}

// ConvertCmd represents the Convert command
var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert OSCAL to CSV or CSV to OSCAL",
	Long:  `Used to convert between data formats`,
}

var ConvertCatalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Convert OSCAL Catalog to CSV or CSV to OSCAL Catalog",
	Long:  `Convert OSCAL Catalog to CSV or CSV to OSCAL Catalog`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Convert Executed")

		// Process whether the combo is permitted
		target, err := processConversionPath(opts.InputFile, opts.OutputFile)
		if err != nil {
			slog.Error("unable to process conversion path", err)
			os.Exit(1)
		}

		if target == "OSCAL" {
			// Convert from CSV to OSCAL
			records, err := common.ReadCSVFile(opts.InputFile)
			if err != nil {
				slog.Error("unable to read csv file", err)
			}

			// Convert to catalog
			catalog, err := oscal.CSVToCatalog(records)
			if err != nil {
				slog.Error("unable to convert to OSCAL", err)
				os.Exit(1)
			}

			err = oscal.WriteCatalog(catalog, opts.OutputFile)
			if err != nil {
				slog.Error("unable to write to file", err)
				os.Exit(1)
			}

		} else {
			// Convert from OSCAL to CSV
			data, err := common.ReadFile(opts.InputFile)
			if err != nil {
				slog.Error("unable to read input file", err)
				os.Exit(1)
			}

			// create catalog model object from []byte
			records, err := oscal.CatalogToCSV(data)
			if err != nil {
				slog.Error("unable to convert to CSV", err)
			}

			err = common.WriteCSV(records, opts.OutputFile)
			if err != nil {
				slog.Error("unable to file", err)
				os.Exit(1)
			}
		}

		slog.Info(fmt.Sprintf("File written to %s\n", opts.OutputFile))

	},
}

func init() {
	rootCmd.AddCommand(ConvertCmd)
	ConvertCmd.AddCommand(ConvertCatalogCmd)

	// Here you will define your flags and configuration settings.
	ConvertCatalogCmd.Flags().StringVarP(&opts.InputFile, "input-file", "f", "", "the path to the input file")
	ConvertCatalogCmd.Flags().StringVarP(&opts.OutputFile, "output-file", "o", "", "the path to write the output file")
	ConvertCatalogCmd.MarkFlagRequired("input-file")
	ConvertCatalogCmd.MarkFlagRequired("output-file")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ConvertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ConvertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func processConversionPath(inputFile string, outputFile string) (string, error) {
	// This should process whether parameter exists and is not empty
	// then should process that the conversion path is permitted
	// then return a string of the conversion path
	if strings.HasSuffix(inputFile, ".csv") {
		if !strings.HasSuffix(outputFile, ".json") && !strings.HasSuffix(outputFile, ".yaml") {
			return "", errors.New("input was a CSV file, output must be json or yaml for OSCAL")
		}
		return "OSCAL", nil
	} else {
		if !strings.HasSuffix(inputFile, ".json") && !strings.HasSuffix(inputFile, ".yaml") {
			return "", errors.New("input was not CSV, must be json or yaml OSCAL")
		}
		if !strings.HasSuffix(outputFile, ".csv") {
			return "", errors.New("output must be a CSV file for conversion from OSCAL")
		}
		return "CSV", nil
	}
}
