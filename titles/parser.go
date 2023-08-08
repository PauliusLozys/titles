package titles

import (
	"errors"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Default configurations.
var (
	DefaultTitleCase    = cases.Title(language.English)
	DefaultTitleRegex   = regexp.MustCompile(`( .+\s?-\s?\d+|.+[sS]\d+([eE]\d{1,2}))|.+[sS]\d+`)
	DefaultSeasonRegex  = regexp.MustCompile(`(([sS]eason\s)|[sS])\d+`)
	DefaultEpisodeRegex = regexp.MustCompile(`(\s?-\s?\d+|[eE]\d+)`)
	DefaultQualityRegex = regexp.MustCompile(`(1080|720|480)p`)
	DefaultReplacer     = strings.NewReplacer(".", " ", "_", " ", "-", " ")
)

type Parser struct {
	titleCase    cases.Caser
	titleRegex   *regexp.Regexp
	seasonRegex  *regexp.Regexp
	episodeRegex *regexp.Regexp
	qualityRegex *regexp.Regexp
	replacer     *strings.Replacer
}

func NewParser(opts ...Option) *Parser {
	tp := &Parser{ // Preset parser with default configuration.
		titleCase:    DefaultTitleCase,
		titleRegex:   DefaultTitleRegex,
		seasonRegex:  DefaultSeasonRegex,
		episodeRegex: DefaultEpisodeRegex,
		qualityRegex: DefaultQualityRegex,
		replacer:     DefaultReplacer,
	}
	for _, opt := range opts {
		opt(tp)
	}
	return tp
}

func (tp *Parser) ParseTitle(unparsedTitle string) (*Title, error) {
	cleanedUpName := cleanUpBrackets(unparsedTitle)
	title := tp.titleRegex.FindString(cleanedUpName)

	// Extract quality.
	quality := tp.qualityRegex.FindString(unparsedTitle)

	// Extract episode number.
	ep := tp.episodeRegex.FindString(title) // Extracted: E01/- 01
	episodeStr := strings.ToLower(ep)
	episodeStr = strings.Trim(episodeStr, " -e")
	episodeNum, err := strconv.Atoi(episodeStr)
	if err != nil { // If no episode was found, default to 1.
		episodeNum = 1
	}

	// Extract season number.
	sea := tp.seasonRegex.FindString(title) // Extracted: S1E01/S1/Season 1
	seasonStr := strings.ToLower(sea)
	if strings.Contains(seasonStr, "season") { // Example: season 1
		seasonStr = strings.TrimSpace(strings.TrimPrefix(seasonStr, "season"))
	} else if strings.TrimSpace(seasonStr) == "" { // no season
		// Titles that don't have a season specified
		// will default to 1. This is usually related to anime episodes.
		seasonStr = "1"

	} else { // Example: s1
		seasonStr = strings.TrimPrefix(seasonStr, "s")
	}
	seasonNum, err := strconv.Atoi(seasonStr)
	if err != nil {
		return nil, err
	}

	// Find where season/episode begins for title cleanup.
	indexOfWhereSeasonEpisodeBegins := -1
	if strings.TrimSpace(sea) != "" {
		indexOfWhereSeasonEpisodeBegins = strings.Index(title, sea)
	}

	if indexOfWhereSeasonEpisodeBegins == -1 {
		ep := tp.episodeRegex.FindStringIndex(title)
		if len(ep) == 0 {
			return nil, errors.New("no seasons/episodes found when parsing title")
		}
		indexOfWhereSeasonEpisodeBegins = ep[0]
	}

	title = tp.replacer.Replace(title[:indexOfWhereSeasonEpisodeBegins])
	cleanedUpTitle := strings.Trim(title, " -]")

	return &Title{
		Name:    tp.titleCase.String(cleanedUpTitle),
		Season:  seasonNum,
		Episode: episodeNum,
		Quality: quality,
	}, nil
}
