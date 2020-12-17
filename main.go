package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"go_news_parser/rss"
	"go_news_parser/search"
)

func main() {
	var err error

	getNews := flag.Bool("get_news", false, "Run get news from lenta.ru or meduza.io")

	searchBy := flag.Bool("search", false, "Search news from DB")
	q := flag.String("q", "", "Search news by title and description fields")
	source := flag.String("source", "", "Search news by a source field")

	flag.Parse()

	if *getNews {
		if len(os.Args) > 2 && len(rss.RssStreams) > 0 {
			err = rss.ParseAndRun(os.Args[2:])
		} else {

			if len(os.Args) <= 1 {
				err = errors.New("Arguments not found")
			}

			if len(rss.RssStreams) == 0 {
				err = errors.New("Rss streams not defined in packeges")
			}
		}
	} else if *searchBy {
		search.SearchNews(*q, *source)
	}

	if err != nil {
		log.Fatal(err)
	}
}
