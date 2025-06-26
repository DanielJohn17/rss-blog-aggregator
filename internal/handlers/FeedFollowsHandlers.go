package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/DanielJohn17/rss-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerCreateFeedFollows(s *State, cmd Command) error {
	cxt := context.Background()

	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: follow <feed_url>")
	}

	feed, err := s.Db.GetFeedByUrl(cxt, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Feed with URL %s does not exist", cmd.Args[0])
	}

	user, err := s.Db.Getuser(cxt, s.Config.CurrentUsername)
	if err != nil {
		return fmt.Errorf("User %s does not exist", s.Config.CurrentUsername)
	}

	follow := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	}

	newFollow, err := s.Db.CreateFeedFollow(cxt, follow)
	if err != nil {
		return fmt.Errorf(
			"Failed to follow feed %s for user %s",
			cmd.Args[0],
			s.Config.CurrentUsername,
		)
	}

	fmt.Printf(
		"Successfully followed feed %s with user %s.\n",
		newFollow.FeedName,
		newFollow.UserName,
	)
	return nil
}

func HandlerGetFeedFollows(s *State, cmd Command) error {
	cxt := context.Background()

	user, err := s.Db.Getuser(cxt, s.Config.CurrentUsername)
	if err != nil {
		return fmt.Errorf("User %s does not exist", s.Config.CurrentUsername)
	}

	follows, err := s.Db.GetFeedFollowsForUser(cxt, user.ID)
	if err != nil {
		return fmt.Errorf("Failed to retrieve feed follows for user %s", s.Config.CurrentUsername)
	}

	if len(follows) == 0 {
		fmt.Println("You are not following any feeds.")
		return nil
	}

	fmt.Printf("Feeds followed by %s:\n", s.Config.CurrentUsername)
	for _, follow := range follows {
		fmt.Printf("- %s\n", follow.FeedName)
	}

	return nil
}
