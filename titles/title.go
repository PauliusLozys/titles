package titles

import "fmt"

type Title struct {
	Name    string // Name if the parsed title.
	Season  int    // Season number.
	Episode int    // Episode number.
}

func (t *Title) String() string {
	return fmt.Sprintf("%s:%d:%d", t.Name, t.Season, t.Episode)
}
