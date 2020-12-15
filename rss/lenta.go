package rss

import "fmt"

type Lenta struct {
	Rss         string
	Rule        string
	ValidRules  []string //{"news", "top7", "last24", "articles"}
	DefaultRule string   //"news"
}

func init() {
	RssStreams = append(RssStreams, "lenta")
}

func (rss *Lenta) Run() {
	fmt.Println("lenta start")
}
