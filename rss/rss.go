package rss

import (
	"errors"
	"fmt"
)

type RssStream interface {
	Run()
	//GetData()
	//ParseDate()
	//Save()
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

func RunParse(runRssStreams []RunRssStream) error {
	var err error
	var rssStream RssStream

	for _, runRssStream := range runRssStreams {
		switch runRssStream.rss {
		case "lenta":
			rssStream = &Lenta{}
		case "meduza":
			rssStream = &Meduza{}
		default:
			err = errors.New(fmt.Sprintf("Couldn't match the %s Rss with the Rss Stream Interface", runRssStream.rss))
		}

		rssStream.Run()
	}

	return err
}
