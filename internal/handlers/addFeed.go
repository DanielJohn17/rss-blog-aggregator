package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielJohn17/rss-blog-aggregator/internal/database"
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
