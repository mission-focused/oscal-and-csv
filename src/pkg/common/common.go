package common

import (
	"encoding/csv"
	"os"
)

func WriteToCSV(rows [][]string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	err = w.WriteAll(rows)

	if err != nil {
		return err
	}
	return nil
}
