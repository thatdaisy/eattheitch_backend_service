package utils

import (
	"fmt"
	"os"
)

func SetIfNotNil[T any](dst *T, src *T) {
	if src != nil {
		*dst = *src
	}
}

func MustGetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("required environment variable %s is not set", key))
	}
	return value
}
