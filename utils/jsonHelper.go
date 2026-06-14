package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Identifiable interface {
	GetID() uuid.UUID
}

var ErrNotFound = errors.New("data: no element with the given id found")

func UpsertJSON[T Identifiable](path string, item T) error {
	items, err := ReadJSON[T](path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	found := false
	for i, existing := range items {
		if existing.GetID() == item.GetID() {
			items[i] = item
			found = true
			break
		}
	}
	if !found {
		items = append(items, item)
	}

	if err := writeJSONArray(path, items); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	return nil
}

func DeleteJSON[T Identifiable](path string, id uuid.UUID) error {
	items, err := ReadJSON[T](path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}

	itemMap := make(map[uuid.UUID]T)
	for _, item := range items {
		itemMap[item.GetID()] = item
	}

	_, exists := itemMap[id]
	if exists {
		delete(itemMap, id)
	}

	items = make([]T, 0, len(itemMap))

	for _, item := range itemMap {
		items = append(items, item)
	}

	if err := writeJSONArray(path, items); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}

	return nil
}

func ReadJSON[T Identifiable](path string) ([]T, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return []T{}, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var items []T
	if err := json.NewDecoder(f).Decode(&items); err != nil {
		return nil, err
	}
	return items, nil
}

func writeJSONArray[T Identifiable](path string, items []T) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(items)
}
