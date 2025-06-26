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

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("The login handler expects a single argument")
	}
	cxt := context.Background()

	user, err := s.Db.Getuser(cxt, cmd.Args[0])
	if err != nil {
		return fmt.Errorf("User %s not found", cmd.Args[0])
	}

	s.Config.SetUser(user.Name)

	fmt.Printf("User %s has been set.\n", cmd.Args[0])

	return nil
}

func HandlerResetUsers(s *State, _ Command) error {
	cxt := context.Background()

	if err := s.Db.Deleteusers(cxt); err != nil {
		return fmt.Errorf("Error deleting users")
	}

	fmt.Println("All users have been reset.")
	return nil
}

func HandlerGetUsers(s *State, _ Command) error {
	cxt := context.Background()

	users, err := s.Db.Getusers(cxt)
	if err != nil {
		return fmt.Errorf("Error retrieving users")
	}

	loggedInUser := s.Config.CurrentUsername

	if users == nil {
		fmt.Println("No users found.")
		return nil
	}

	for _, u := range users {
		if u.Name == loggedInUser {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}

	return nil
}
