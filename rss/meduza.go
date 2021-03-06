package rss

import (
	"encoding/json"
	"fmt"
	"go_news_parser/news"
	"go_news_parser/restclient"
	"log"
	"strconv"
	"time"
)

type MeduzaCatalog struct {
	Documents map[string]MeduzaDocument `json:"documents"`
	Count     int                       `json:"_count"`
}

type MeduzaDocument struct {
	Link        string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"second_title"`
	Date        int64  `json:"published_at"`
}

// The Meduza struct contain functions for getting news list
type Meduza struct {
	Rule uint
}

func init() {
	RssStreams = append(RssStreams, "meduza")
}

// Initialization function checks injected a rule. If a rule is not defined, will be add default rule.
func (rss *Meduza) Initialization(ruleArgument string) {

	if ruleArgument == "" {
		ruleArgument = "0"
	}

	rule, err := strconv.ParseUint(ruleArgument, 10, 32)

	if err != nil {
		log.Printf("Meduza rule: %s: %v\n", ruleArgument, err)

		rule = 0
	}

	rss.Rule = uint(rule)

}

// GetNewsList function returns news list from Meduza rss stream
func (rss *Meduza) GetNewsList() error {
	var catalog MeduzaCatalog
	var response []byte
	var err error

	response, err = restclient.Request(fmt.Sprintf("http://meduza.io/api/v3/search?chrono=news&locale=ru&per_page=24&page=%d", rss.Rule))
	if err != nil {

		return err
	}

	err = json.Unmarshal(response, &catalog)
	if err == nil {
		if catalog.Count != 0 {
			for _, item := range catalog.Documents {

				newsItem := news.News{
					Link:        fmt.Sprintf("https://meduza.io/%s", item.Link),
					Title:       item.Title,
					Description: item.Description,
					Source:      "meduza",
					DatePub:     time.Unix(item.Date, 0),
				}

				SaveNews(newsItem)
			}
		} else {
			log.Println("Meduza rss is empty")
		}
	}

	return err
}

// GetRule function return parsing rule
func (rss *Meduza) GetRule() string {

	return fmt.Sprintf("%d", rss.Rule)
}
