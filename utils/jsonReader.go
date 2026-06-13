package utils

import (
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
