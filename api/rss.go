package api

type RSSFeed struct {
	Channel struct {
		Title       string
		Link        string
		Description string
		Language    string
		Item        []RSSItem
	}
}

type RSSItem struct {
	Title       string
	Link        string
	Description string
	PubDate     string
}
