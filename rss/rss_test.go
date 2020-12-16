package rss

import (
	"go_news_parser/news"
	"log"
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

func TestLentaRss(t *testing.T) {
	var rssStreamObject RssStream
	var newsList []news.News
	var err error
	var hasError bool

	rules := []string{"news", "top7", "last24", "articles"}

	rssStreamObject = &Lenta{}

	for _, rule := range rules {
		rssStreamObject.Initialization(rule)
		newsList, err = rssStreamObject.GetNewsList()

		log.Printf("Lenta %s - %d\n", rule, len(newsList))

		if err != nil {
			hasError = true
			break
		}
	}

	if hasError {
		t.Error(err)
	}
}

func TestMeduzaRss(t *testing.T) {
	var rssStreamObject RssStream
	var newsList []news.News
	var err error
	var hasError bool

	rules := []string{"1", "r", "-4", "0"}

	rssStreamObject = &Meduza{}

	for _, rule := range rules {
		rssStreamObject.Initialization(rule)
		newsList, err = rssStreamObject.GetNewsList()

		log.Printf("Meduza %s - %d\n", rule, len(newsList))

		if err != nil {
			hasError = true
			break
		}
	}

	if hasError {
		t.Error(err)
	}
}
