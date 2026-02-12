package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type File struct {
	ParsedShowName   string
	UnparsedShowName string
	FilePath         string
	Season           int
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
			UnparsedShowName: file.Name(),
			FilePath:         filepath.Join(startDir, file.Name()),
		})
	}
	return list
}

func moveFiles(list []File) {
	organized := make(map[string][]File, len(list))
	for _, file := range list {
		organized[file.ParsedShowName] = append(organized[file.ParsedShowName], file)
	}

	for showName, episodes := range organized {
		var foundExistingFolder bool
		if *matchExistingFolder {
			// Try finding an existing folder with similar name (case insensitive).
			// Avoid creating duplicate folders if something like folder case is different.
			showName, foundExistingFolder = findFolderIfExists(*outputDir, showName)
		}

		showFolder := filepath.Join(*outputDir, showName)
		for _, episode := range episodes {
			if episode.Season == 0 { // Assume file was unparsed.
				fmt.Println("ERROR: Unparsed file, skipping:", episode.UnparsedShowName)
				continue
			}

			seasonFolder := filepath.Join(showFolder, fmt.Sprintf("Season %d", episode.Season))
			createFolderIfNeeded(seasonFolder, *dryRun) // create folder for show/season, if needed
			finalPath := filepath.Join(seasonFolder, episode.UnparsedShowName)

			if !*dryRun {
				if err := os.Rename(episode.FilePath, finalPath); err != nil {
					fmt.Println("ERROR: moving files:", err)
					continue
				}
			}

			if foundExistingFolder {
				fmt.Println("Moved to existing folder:", finalPath)
			} else {
				fmt.Println("Moved to a new folder:", finalPath)
			}
		}
	}
}
