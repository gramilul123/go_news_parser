package rss

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"go_news_parser/news"
	"go_news_parser/restclient"
)

type LentaCatalog struct {
	Items LentaChannel `xml:"channel"`
}

type LentaChannel struct {
	Items []LentaItem `xml:"item"`
}

type LentaItem struct {
	Link        string `xml:"guid"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Date        string `xml:"pubDate"`
}

// The Lenta struct contain functions for getting news list
type Lenta struct {
	Rule string
}

func init() {
	RssStreams = append(RssStreams, "lenta")
}

// Initialization function checks injected a rule. If a rule is not defined, will be add default rule.
func (rss *Lenta) Initialization(ruleArgument string) {
	validRules := []string{"news", "top7", "last24", "articles"}
	defaultRule := "news"

	for _, rule := range validRules {
		if ruleArgument == rule {
			rss.Rule = rule

			break
		}
	}

	if rss.Rule == "" {
		log.Printf("For Lenta rss will be use defaul rule %s\n", defaultRule)
		rss.Rule = defaultRule
	}
}

// GetNewsList function returns news list from Lenta rss stream
func (rss *Lenta) GetNewsList() error {
	var catalog LentaCatalog
	var response []byte
	var err error
	var date time.Time

	response, err = restclient.Request(fmt.Sprintf("http://lenta.ru/rss/%s", rss.Rule))
	if err != nil {

		return err
	}

	err = xml.Unmarshal(response, &catalog)

	if err == nil {
		if len(catalog.Items.Items) > 0 {
			for _, item := range catalog.Items.Items {

				date, err = time.Parse(time.RFC1123Z, item.Date)

				if err == nil {

					newsItem := news.News{
						Link:        item.Link,
						Title:       item.Title,
						Description: item.Description,
						Source:      "lenta",
						DatePub:     date,
					}

					SaveNews(newsItem)
				}
			}
		} else {
			log.Println("Lenta rss is empty")
		}
	}

	return err
}

// GetRule function return parsing rule
func (rss *Lenta) GetRule() string {

	return rss.Rule
}
