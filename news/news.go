package news

import "time"

type News struct {
	ID          int    `gorm:"primaryKey"`
	Link        string `gorm:"primaryKey"`
	Title       string
	Description string
	Source      string
	DatePub     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
