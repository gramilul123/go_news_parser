package main

import (
	"errors"
	"log"
	"os"

	"go_news_parser/rss"
)

func main() {
	var err error

	if len(os.Args) > 1 && len(rss.RssStreams) > 0 {
		err = rss.ParseAndRun(os.Args[1:])
	} else {

		if len(os.Args) <= 1 {
			err = errors.New("Arguments not found")
		}

		if len(rss.RssStreams) == 0 {
			err = errors.New("Rss streams not defined in packeges")
		}

	}

	if err != nil {
		log.Fatal(err)
	}
}
