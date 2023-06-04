package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		unparsedName   string
		expectedTitle  string
		expectedSeason int
		expectedErr    error
	}{
		{
			unparsedName:   "The Office - S01E01 - Pilot.mkv",
			expectedTitle:  "The Office",
			expectedSeason: 1,
			expectedErr:    nil,
		},
		{
			unparsedName:   "[Anime Time] Tengoku Daimakyou - 10 [1080p][HEVC 10bit x265][AAC][Multi Sub]",
			expectedTitle:  "Tengoku Daimakyou",
			expectedSeason: 1,
			expectedErr:    nil,
		},
		{
			unparsedName:   "[Trix] Heavenly Delusion - S01E07 - (1080p AV1 AAC)[Multi Subs]",
			expectedTitle:  "Heavenly Delusion",
			expectedSeason: 1,
			expectedErr:    nil,
		},
		{
			unparsedName:   "Doom.Patrol.S02E05.1080p.HEVC.x265-MeGusta[eztv.re].mkv",
			expectedTitle:  "Doom Patrol",
			expectedSeason: 2,
			expectedErr:    nil,
		},
		{
			unparsedName:   "[Erai-raws] Isekai Ojisan - 12 [1080p][Multiple Subtitle][18176279].mkv",
			expectedTitle:  "Isekai Ojisan",
			expectedSeason: 1,
			expectedErr:    nil,
		},
		{
			unparsedName:   "The.Flash.2014.S09E11.1080p.HEVC.x265-MeGusta[eztv.re].mkv",
			expectedTitle:  "The Flash 2014",
			expectedSeason: 9,
			expectedErr:    nil,
		},
		{
			unparsedName:   "Doom.Patrol.1080p.HEVC.x265-MeGusta[eztv.re].mkv",
			expectedTitle:  "",
			expectedSeason: -1,
			expectedErr:    errors.New("no seasons/episodes found in anime title string"),
		},
	}

	for _, test := range tests {
		actualTitle, actualSeason, actualErr := parseFile(test.unparsedName)
		assert.Equal(t, test.expectedTitle, actualTitle)
		assert.Equal(t, test.expectedSeason, actualSeason)
		assert.Equal(t, test.expectedErr, actualErr)
	}
}
