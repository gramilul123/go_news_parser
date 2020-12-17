package search

import (
	"fmt"

	"go_news_parser/db"
	"go_news_parser/news"
)

func SearchNews(q, source string) {
	newsList := GetNews(q, source)

	if len(newsList) > 0 {

		for _, item := range newsList {
			fmt.Printf("Source: %s; Link: %s; Title: %v; Date: %v\n", item.Source, item.Link, item.Title, item.DatePub)
		}
	} else {
		fmt.Println("News not found")
	}
}

func GetNews(q, source string) []news.News {
	var newsList []news.News

	db := db.GetDB()
	chain := db.Where("")

	if source != "" {
		chain = db.Where("source = ?", source)
	}

	if q != "" {
		chain = db.Where("(title LIKE ? OR description LIKE ?)", fmt.Sprintf("%%%s%%", q), fmt.Sprintf("%%%s%%", q))
	}

	chain.Order("date_pub desc").Find(&newsList)

	return newsList
}
