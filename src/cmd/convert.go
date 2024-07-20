/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ConvertCmd represents the Convert command
var ConvertCmd = &cobra.Command{
	Use:   "Convert",
	Short: "Convert OSCAL to CSV or CSV to OSCAL",
	Long:  `Used to convert between data formats`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Convert called")
	},
}

func init() {
	rootCmd.AddCommand(ConvertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ConvertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ConvertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
