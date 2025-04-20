package titles

import "strings"

const maxIterations = 50

func cleanUpBrackets(fileName string) string {
	for range maxIterations {
		s := strings.IndexRune(fileName, '[')
		f := strings.IndexRune(fileName, ']')
		if s == -1 || f == -1 || s > f {
			return fileName
		}
		fileName = fileName[:s] + fileName[f+1:]
	}

	panic("unreachable")
}
