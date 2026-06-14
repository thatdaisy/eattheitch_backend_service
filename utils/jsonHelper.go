package utils

import (
	"encoding/json"
	"errors"
	"os"
)

func ReadJson(filepath string) ([]byte, error) {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("file empty")
	}
	return data, nil
}

func WriteJson[T any](filepath string, items []T) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath, data, 0644)
}
