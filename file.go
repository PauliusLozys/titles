package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const maxIterations = 50

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

func cleanUpBrackets(fileName string) string {
	for i := 0; i < maxIterations; i++ {
		s := strings.IndexRune(fileName, '[')
		f := strings.IndexRune(fileName, ']')
		if s == -1 || f == -1 || s > f {
			return fileName
		}
		fileName = fileName[:s] + fileName[f+1:]
	}
	return fileName // should be unreachable
}

func parseFile(unparsedName string) (string, int, error) {
	cleanedUpName := cleanUpBrackets(unparsedName)
	rawName := cleanedUpName
	title := title.FindString(rawName)
	sea := season.FindString(title) // Extracted: S1E01/S1/Season 1
	seaNum := strings.ToLower(sea)

	if strings.Contains(seaNum, "season") { // Season 1
		seaNum = strings.TrimSpace(strings.TrimPrefix(seaNum, "season"))

	} else if strings.TrimSpace(seaNum) == "" { // no season
		// Episodes that don't have a season specified
		// will default to 1. This is usually related to anime episodes.
		seaNum = "1"

	} else { // S1
		seaNum = seaNum[1:]
	}

	seasonNum, err := strconv.Atoi(seaNum)
	if err != nil {
		return "", -1, err
	}

	lastIndex := -1
	if strings.TrimSpace(sea) != "" {
		lastIndex = strings.Index(title, sea)
	}

	if lastIndex == -1 {
		ep := episode.FindStringIndex(title)
		if len(ep) == 0 {
			return "", -1, errors.New("no seasons/episodes found in anime title string")
		}
		lastIndex = ep[0]
	}
	title = replacer.Replace(title[:lastIndex])
	return strings.Trim(title, " -]"), seasonNum, nil
}
