package titles

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
)

type Option func(*Parser)

func WithTitleTransformer(titleCase cases.Caser) Option {
	return func(tp *Parser) {
		tp.titleCase = titleCase
	}
}

func WithTitleRegex(titleRegex *regexp.Regexp) Option {
	return func(tp *Parser) {
		tp.titleRegex = titleRegex
	}
}

func WithSeasonRegex(seasonRegex *regexp.Regexp) Option {
	return func(tp *Parser) {
		tp.seasonRegex = seasonRegex
	}
}

func WithEpisodeRegex(episodeRegex *regexp.Regexp) Option {
	return func(tp *Parser) {
		tp.episodeRegex = episodeRegex
	}
}

func WithQualityRegex(qualityRegex *regexp.Regexp) Option {
	return func(tp *Parser) {
		tp.qualityRegex = qualityRegex
	}
}

func WithReplacer(replacer *strings.Replacer) Option {
	return func(tp *Parser) {
		tp.replacer = replacer
	}
}
