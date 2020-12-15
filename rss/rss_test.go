package rss

import (
	"strings"
	"testing"
)

type lineArguments struct {
	line  string
	count int
}

func TestParseLineArguments(t *testing.T) {
	errorLines := []int{}

	lineArguments := []lineArguments{
		{"lenta top7 meduza 1-7 meduza 2 lenta news", 4},
		{"lenta", 1},
		{"any all", 0},
		{"lenta meduza 10", 2},
		{"lenta meduza", 2},
		{"any lenta", 1},
	}

	for key, line := range lineArguments {
		args := strings.Split(line.line, " ")
		rssStreams := ParseLineArguments(args)

		if len(rssStreams) != line.count {
			errorLines = append(errorLines, key)
		}
	}

	if len(errorLines) > 0 {
		for _, errorLine := range errorLines {
			t.Errorf("Error line: %d (%v)", errorLines, lineArguments[errorLine])
		}
	}
}
