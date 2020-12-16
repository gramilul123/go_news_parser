package news

import "time"

type News struct {
	ID          int
	Link        string
	Title       string
	Description string
	Source      string
	DatePub     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
