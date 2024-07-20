/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// GenerateCmd represents the Generate command
var GenerateCmd = &cobra.Command{
	Use:   "Generate",
	Short: "Generates a CSV template for a specified OSCAL model",
	Long: `Generates a CSV template for a specified OSCAL model
	Examples:
		oac generate catalog`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Generate called")
	},
}

func init() {
	rootCmd.AddCommand(GenerateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// GenerateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// GenerateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
