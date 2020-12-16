package rss

import (
	"encoding/json"
	"fmt"
	"go_news_parser/news"
	"go_news_parser/restclient"
	"log"
	"strconv"
)

type MeduzaCatalog struct {
	Documents map[string]MeduzaDocument `json:"documents"`
	Count     int                       `json:"_count"`
}

type MeduzaDocument struct {
	Link        string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"second_title"`
	Date        int    `json:"published_at"`
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

	rule, err := strconv.ParseUint(ruleArgument, 10, 32)

	if err != nil {
		log.Printf("Meduza rule: %s: %v\n", ruleArgument, err)

		rule = 0
	}

	rss.Rule = uint(rule)

}

// GetNewsList function returns news list from Meduza rss stream
func (rss *Meduza) GetNewsList() ([]news.News, error) {
	var catalog MeduzaCatalog
	var newsList []news.News
	var response []byte
	var err error

	response, err = restclient.Request(fmt.Sprintf("http://meduza.io/api/v3/search?chrono=news&locale=ru&per_page=24&page=%d", rss.Rule))
	if err != nil {

		return nil, err
	}

	err = json.Unmarshal(response, &catalog)
	if err == nil {
		if catalog.Count != 0 {
			for _, item := range catalog.Documents {
				newsList = append(newsList, news.News{0, item.Link, item.Title, item.Description, "Meduza", fmt.Sprint(item.Date)})
			}
		} else {
			log.Println("Meduza rss is empty")
		}
	}

	return newsList, err
}
