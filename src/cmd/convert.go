/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mission-focused/oscal-and-csv/src/pkg/common"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
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
		fmt.Println("Convert Catalog called")

		// Process whether the combo is permitted
		target, err := processConversionPath(opts.InputFile, opts.OutputFile)
		if err != nil {
			slog.Error("unable to process conversion path", err)
			os.Exit(1)
		}

		data, err := common.ReadFile(opts.InputFile)
		if err != nil {
			slog.Error("unable to read input file", err)
			os.Exit(1)
		}

		if target == "OSCAL" {
			// Convert from CSV to OSCAL

		} else {
			// Convert from OSCAL to CSV
		}

	},
}

func init() {
	rootCmd.AddCommand(ConvertCmd)

	// Here you will define your flags and configuration settings.
	ConvertCatalogCmd.Flags().StringVarP(&opts.InputFile, "input-file", "i", "", "the path to the input file")
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
