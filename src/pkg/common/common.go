package common

import (
	"encoding/csv"
	"errors"
	"fmt"
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
