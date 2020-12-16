package news

import "time"

type News struct {
	ID          int
	Link        string `sql:"unique_index:idx_link"`
	Title       string
	Description string
	Source      string
	DatePub     time.Time
}
