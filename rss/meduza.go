package rss

type Rules struct {
	Page    uint
	PerPage uint
}

type Meduza struct {
	Rss  string
	Rule Rules
}

func init() {
	RssStreams = append(RssStreams, "meduza")
}
