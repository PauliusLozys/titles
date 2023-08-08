package titles_test

import (
	"errors"
	"testing"

	"github.com/PauliusLozys/titles/titles"
	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		unparsedTitle string
		expected      *titles.Title
		err           error
	}{
		{
			unparsedTitle: "The Office - S01E01 - Pilot.mkv",
			expected:      &titles.Title{Name: "The Office", Season: 1, Episode: 1},
		},
		{
			unparsedTitle: "[Anime Time] Tengoku Daimakyou - 10 [1080p][HEVC 10bit x265][AAC][Multi Sub]",
			expected:      &titles.Title{Name: "Tengoku Daimakyou", Season: 1, Episode: 10},
		},
		{
			unparsedTitle: "[Trix] Heavenly Delusion - S01E07 - (1080p AV1 AAC)[Multi Subs]",
			expected:      &titles.Title{Name: "Heavenly Delusion", Season: 1, Episode: 7},
		},
		{
			unparsedTitle: "Doom.Patrol.S02E05.1080p.HEVC.x265-MeGusta[eztv.re].mkv",
			expected:      &titles.Title{Name: "Doom Patrol", Season: 2, Episode: 5},
		},
		{
			unparsedTitle: "[Erai-raws] Isekai Ojisan - 12 [1080p][Multiple Subtitle][18176279].mkv",
			expected:      &titles.Title{Name: "Isekai Ojisan", Season: 1, Episode: 12},
		},
		{
			unparsedTitle: "The.Flash.2014.S09E11.1080p.HEVC.x265-MeGusta[eztv.re].mkv",
			expected:      &titles.Title{Name: "The Flash 2014", Season: 9, Episode: 11},
		},
		{
			unparsedTitle: "Doom.Patrol.1080p.HEVC.x265-MeGusta[eztv.re].mkv",
			expected:      nil,
			err:           errors.New("no seasons/episodes found when parsing title"),
		},
	}

	parser := titles.NewParser()
	for _, test := range tests {
		t.Run(test.unparsedTitle, func(t *testing.T) {
			actual, err := parser.ParseTitle(test.unparsedTitle)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.err, err)
		})
	}
}
