package rss

type Strategy interface {
	GetData()
	ParseDate()
	Save()
}

// ParseAndRun function parses arguments and run rss parser
func ParseAndRun(args []string) error {
	var err error

	return err
}
