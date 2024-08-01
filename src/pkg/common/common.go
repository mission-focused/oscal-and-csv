package common

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

func WriteCSV(rows [][]string, filePath string) error {
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

func ReadCSVFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return [][]string{}, fmt.Errorf("unable to read input file: %w", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Errorf("unable to parse file as CSV for %s with error: %w", filePath, err)
	}

	return records, nil
}

func CheckFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !errors.Is(err, os.ErrNotExist)
}

func ReadFile(filePath string) ([]byte, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return []byte{}, fmt.Errorf("path: %v does not exist - unable to digest document", filePath)
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return data, err
	}
	return data, nil
}
