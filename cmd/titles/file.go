package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type File struct {
	ParsedName   string
	UnparsedName string
	FilePath     string
	Season       int
}

func collectAllFiles(startDir string, blacklistedDirs, extensions []string, recursive bool) []File {
	blacklist := toMap(blacklistedDirs)
	extension := toMap(extensions)

	files, err := os.ReadDir(startDir)
	panicOnError(err)
	list := make([]File, 0)
	for _, file := range files {
		if file.IsDir() && recursive {
			if _, ok := blacklist[file.Name()]; ok {
				continue
			}
			list = append(list, collectAllFiles(filepath.Join(startDir, file.Name()), blacklistedDirs, extensions, recursive)...)
			continue
		}
		ext := path.Ext(file.Name())
		_, ok := extension[ext]
		if strings.HasPrefix(file.Name(), ".") || !ok {
			continue
		}
		list = append(list, File{
			UnparsedName: file.Name(),
			FilePath:     filepath.Join(startDir, file.Name()),
		})
	}
	return list
}

func moveFiles(list []File, outputDir string, dryRun bool) {
	organized := make(map[string][]File, len(list))
	for _, file := range list {
		organized[file.ParsedName] = append(organized[file.ParsedName], file)
	}

	for name, files := range organized {
		showFolder := filepath.Join(outputDir, name)
		createFolderIfNeeded(showFolder, dryRun) // create folder for show, if needed
		for _, file := range files {
			if file.Season == 0 { // Assume file was unparsed.
				continue
			}
			season := fmt.Sprintf("Season %d", file.Season)
			seasonFolder := filepath.Join(showFolder, season)
			createFolderIfNeeded(seasonFolder, dryRun) // create folder for season, if needed
			finalPath := filepath.Join(seasonFolder, file.UnparsedName)
			if !dryRun {
				if err := os.Rename(file.FilePath, finalPath); err != nil {
					fmt.Println("ERROR: moving files:", err)
					continue
				}
			}
			fmt.Println("Moved to:", finalPath)
		}
	}
}
