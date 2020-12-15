package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"go_news_parser/rss"
)

func main() {
	var err error

	if len(os.Args) > 1 && len(rss.Streams) > 0 {
		err = rss.ParseAndRun(os.Args[1:])
	} else {

		if len(os.Args) <= 1 {
			err = errors.New("Arguments not found")
		}

		if len(rss.Streams) == 0 {
			err = errors.New("Rss streams not defined in packeges")
		}

	}

	fmt.Println(rss.Streams)

	if err != nil {
		log.Fatal(err)
	}
}
