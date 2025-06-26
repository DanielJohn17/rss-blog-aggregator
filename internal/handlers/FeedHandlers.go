package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielJohn17/rss-blog-aggregator/internal/database"
	"github.com/DanielJohn17/rss-blog-aggregator/internal/feed"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	cxt := context.Background()

	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}

	user, err := s.Db.Getuser(cxt, s.Config.CurrentUsername)
	if err != nil {
		return fmt.Errorf("User %s does not exist", s.Config.CurrentUsername)
	}

	feed := database.CreatefeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    user.ID,
	}

	newFeed, err := s.Db.Createfeed(cxt, feed)
	if err != nil {
		return fmt.Errorf(
			"Feed %s already exists for user %s",
			cmd.Args[0],
			s.Config.CurrentUsername,
		)
	}

	fmt.Printf(
		"Feed %s has been created with ID %s for user %s.\n",
		newFeed.Name,
		newFeed.ID.String(),
		s.Config.CurrentUsername,
	)
	return nil
}

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

func HandlerGetFeeds(s *State, cmd Command) error {
	ctx := context.Background()

	feeds, err := s.Db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error fetching feeds")
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("User Name: %s\n", feed.UserName)
		fmt.Printf("Created At: %s\n", feed.CreatedAt.Format("01-02-25"))
		fmt.Printf("Updated At: %s\n", feed.UpdatedAt.Format("01-02-25"))
		fmt.Println("-----------------------------")
	}

	return nil
}
