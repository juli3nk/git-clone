package main

import (
	"os"
	"path"
	"strings"
)

func readSecret(key string) (string, error) {
	basePath := "/run/secrets"

	readPath := path.Join(basePath, key)
	secretBytes, err := os.ReadFile(readPath)
	if err != nil {
		return "", err
	}

	val := strings.TrimSpace(string(secretBytes))

	return val, nil
}
