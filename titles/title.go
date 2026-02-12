package titles

import "fmt"

type Title struct {
	Name    string // Name of the parsed title.
	Season  int    // Season number.
	Episode int    // Episode number.
	Quality string // Quality of the video. Example: 1080p, 720p, etc.
}

func (t *Title) String() string {
	return fmt.Sprintf("%s:%d:%d:%s", t.Name, t.Season, t.Episode, t.Quality)
}
