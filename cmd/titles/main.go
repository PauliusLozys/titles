package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PauliusLozys/titles/titles"
)

var (
	baseDir         = flag.String("i", ".", "input directory")
	outputDir       = flag.String("o", "./output", "output directory")
	blacklistedDirs = flag.String("b", "", "blacklisted directories separated by ','. Example: './dir1,./dir2'")
	extensions      = flag.String("e", ".mkv,.mp4", "file extension to look for separated by ','")
	recursive       = flag.Bool("r", false, "recursively search for all files")
	dryRun          = flag.Bool("d", false, "do a dry run without affecting files")
)

func main() {
	flag.Parse()

	list := collectAllFiles(
		*baseDir,
		strings.Split(*blacklistedDirs, ","),
		strings.Split(*extensions, ","),
		*recursive,
	)

	parser := titles.NewParser()

	for i := range list {
		if !titles.DefaultTitleRegex.MatchString(list[i].UnparsedName) {
			fmt.Println("Unmatched file:", list[i].UnparsedName)
			continue
		}

		title, err := parser.ParseTitle(list[i].UnparsedName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		list[i].ParsedName = title.Name
		list[i].Season = title.Season
	}
	moveFiles(list, *outputDir, *dryRun)
}
