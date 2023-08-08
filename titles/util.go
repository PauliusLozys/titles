package titles

import "strings"

const maxIterations = 50

func cleanUpBrackets(fileName string) string {
	for i := 0; i < maxIterations; i++ {
		s := strings.IndexRune(fileName, '[')
		f := strings.IndexRune(fileName, ']')
		if s == -1 || f == -1 || s > f {
			return fileName
		}
		fileName = fileName[:s] + fileName[f+1:]
	}
	return fileName // should be unreachable
}
