package handlers

import (
	"context"
	"fmt"
)

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
