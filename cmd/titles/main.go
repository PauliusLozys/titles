package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/PauliusLozys/titles/titles"
)

var (
	baseDir             = flag.String("i", ".", "input directory")
	outputDir           = flag.String("o", "./output", "output directory")
	blacklistedDirs     = flag.String("b", "", "blacklisted directories separated by ','. Example: './dir1,./dir2'")
	extensions          = flag.String("e", ".mkv,.mp4", "file extension to look for separated by ','")
	recursive           = flag.Bool("r", false, "recursively search for all files")
	dryRun              = flag.Bool("d", false, "do a dry run without affecting files")
	matchExistingFolder = flag.Bool("m", true, "try to match already existing folder in the output directory using case-insensitive regex")
	version             = flag.Bool("v", false, "print version and exit")

	COMMIT     = "" // Set at build time using -ldflags "-X main.COMMIT=$(git rev-parse HEAD)"
	BUILD_TIME = "" // Set at build time using -ldflags "-X main.BUILD_TIME=$(date -u +%Y-%m-%dT%H:%M:%SZ)"
)

func main() {
	flag.Parse()

	if *version {
		fmt.Printf("Commit: %s\nBuild Time: %s\n", COMMIT, BUILD_TIME)
		return
	}

	list := collectAllFiles(
		*baseDir,
		strings.Split(*blacklistedDirs, ","),
		strings.Split(*extensions, ","),
		*recursive,
	)

	if len(list) == 0 {
		slog.Warn("No files were found")
		os.Exit(0)
	}

	parser := titles.NewParser()

	for i := range list {
		if !titles.DefaultTitleRegex.MatchString(list[i].UnparsedShowName) {
			slog.Warn("Unmatched file", slog.String("file", list[i].UnparsedShowName))

			continue
		}

		title, err := parser.ParseTitle(list[i].UnparsedShowName)
		if err != nil {
			slog.Error(
				"Parsing title",
				slog.String("title", list[i].UnparsedShowName),
				slog.Any("err", err),
			)

			continue
		}

		list[i].ParsedShowName = title.Name
		list[i].Season = title.Season
	}

	moveFiles(list)
}
