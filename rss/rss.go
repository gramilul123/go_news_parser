package rss

import (
	"errors"
	"fmt"
	"go_news_parser/db"
	"go_news_parser/news"
	"log"

	"gorm.io/gorm/clause"
)

type RssStream interface {
	Initialization(string)
	GetNewsList() ([]news.News, error)
}

type RunRssStream struct {
	rss  string
	rule string
}

var RssStreams []string

// ParseAndRun function parses arguments and run rss parser
func ParseAndRun(args []string) error {
	var err error
	runRssStreams := ParseLineArguments(args)

	if len(runRssStreams) > 0 {

		err = RunParse(runRssStreams)
	} else {

		err = errors.New("Couldn't find a command to run the parser")
	}

	return err
}

// ParseLineArguments function parses command line arguments and return rss streams with rule
func ParseLineArguments(args []string) []RunRssStream {
	runRssStreams := []RunRssStream{}

	for key, arg := range args {

		for _, stream := range RssStreams {

			if arg == stream {
				rss := stream

				rule := ""
				if key+1 < len(args) {
					rule = args[key+1]
				}

				runRssStreams = append(runRssStreams, RunRssStream{rss, rule})
			}
		}
	}

	return runRssStreams
}

// RunParse function run parse
func RunParse(runRssStreams []RunRssStream) error {
	var err error
	var rssStreamObject RssStream

	for _, runRssStream := range runRssStreams {
		switch runRssStream.rss {
		case "lenta":
			rssStreamObject = &Lenta{}
		case "meduza":
			rssStreamObject = &Meduza{}
		default:
			err = errors.New(fmt.Sprintf("Couldn't match the %s Rss with the Rss Stream Interface", runRssStream.rss))
		}

		if err == nil {
			err = GetAndSaveNews(rssStreamObject, runRssStream.rss, runRssStream.rule)
		}

		if err != nil {
			break
		}

	}

	return err
}

// GetAndSaveNews function manages rss processing
func GetAndSaveNews(rssStreamObject RssStream, rss, rule string) error {
	var newsList []news.News
	var err error

	rssStreamObject.Initialization(rule)
	newsList, err = rssStreamObject.GetNewsList()

	if err != nil {

		return err
	}

	Save(newsList, rss, rule)

	return err
}

// Save function save news list to DB
func Save(newsList []news.News, rss, rule string) {

	for _, news := range newsList {

		db.GetDB().Clauses(clause.OnConflict{DoNothing: true}).Create(&news)
	}

	log.Printf("Done %s %s\n", rss, rule)
}
