package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	baseDir         = flag.String("i", ".", "input directory")
	outputDir       = flag.String("o", "./output", "output directory")
	blacklistedDirs = flag.String("b", "", "blacklisted directories separated by ','. Example: './dir1,./dir2'")
	extensions      = flag.String("e", ".mkv,.mp4", "file extension to look for separated by ','")
	recursive       = flag.Bool("r", false, "recursively search for all files")
	dryRun          = flag.Bool("d", false, "do a dry run without affecting files")

	tt       = cases.Title(language.English)
	title    = regexp.MustCompile(`( .+\s?-\s?\d+|.+[sS]\d+([eE]\d{1,2}))`)
	season   = regexp.MustCompile(`(([sS]eason\s)|[sS])\d+`)
	episode  = regexp.MustCompile(`(\s?-\s?\d+|[eE]\d+)`)
	replacer = strings.NewReplacer(".", " ", "_", " ", "-", " ")
)

func main() {
	flag.Parse()

	list := collectAllFiles(
		*baseDir,
		strings.Split(*blacklistedDirs, ","),
		strings.Split(*extensions, ","),
		*recursive,
	)

	for i := range list {
		if title.MatchString(list[i].UnparsedName) {
			title, err := parseFile(&list[i])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}
			list[i].ParsedName = tt.String(title)
			continue
		}
		fmt.Println("Unmatched file:", list[i].UnparsedName)
	}
	moveFiles(list, *outputDir, *dryRun)
}
