package main

import (
	"log/slog"
	"os"
	"regexp"
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

func findFolderIfExists(outputDir, name string) string {
	dirs, err := os.ReadDir(outputDir)
	if err != nil {
		slog.Error("could not read output directory", slog.String("dir", outputDir), slog.Any("err", err))
		return name
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		matched, err := regexp.MatchString("(?i)"+dir.Name(), name)
		if err != nil {
			slog.Error("matching directory", slog.String("pattern", name), slog.String("matching string", dir.Name()), slog.Any("err", err))
			return name
		}

		if matched {
			return dir.Name()
		}
	}

	return name
}
