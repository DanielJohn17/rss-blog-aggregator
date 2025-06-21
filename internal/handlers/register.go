package handlers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DanielJohn17/rss-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("The register handler expects a single argument")
	}
	cxt := context.Background()
	user := database.CreateuserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
	}

	newUser, err := s.Db.Createuser(cxt, user)
	if err != nil {
		return fmt.Errorf("User %s already exists", cmd.Args[0])
	}

	HandlerLogin(s, cmd)
	fmt.Printf("User %s has been created with ID %s.\n", newUser.Name, newUser.ID.String())
	return nil
}
