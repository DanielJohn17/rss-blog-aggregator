package handlers

import (
	"context"
	"fmt"

	"github.com/DanielJohn17/rss-blog-aggregator/internal/feed"
)

func HandlerAgg(s *State, cmd Command) error {
	url := "https://www.wagslane.dev/index.xml"
	cxt := context.Background()

	feed, err := feed.FetchFeed(cxt, url)
	if err != nil {
		return err
	}

	fmt.Printf("Feed Title: %s\n", feed.Channel.Title)
	fmt.Printf("Description: %s\n", feed.Channel.Description)

	for _, item := range feed.Channel.Item {
		fmt.Printf("Item Title: %s\n", item.Title)
		fmt.Printf("Item Link: %s\n", item.Link)
		fmt.Printf("Item Description: %s\n", item.Description)
		fmt.Printf("Publication Date: %s\n", item.PubDate)
	}

	return nil
}
