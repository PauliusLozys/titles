package main

import (
	"os"
)

func toMap(slice []string) map[string]struct{} {
	m := make(map[string]struct{}, len(slice))
	for _, v := range slice {
		m[v] = struct{}{}
	}
	return m
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func createFolderIfNeeded(path string, dryRun bool) {
	if dryRun {
		return
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// No folder
		panicOnError(os.MkdirAll(path, 0755))
	}
}
