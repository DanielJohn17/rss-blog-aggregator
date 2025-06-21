package handlers

import (
	"context"
	"fmt"
)

func HandlerGetUsers(s *State, _ Command) error {
	cxt := context.Background()

	users, err := s.Db.Getusers(cxt)
	if err != nil {
		return fmt.Errorf("Error retrieving users")
	}

	loggedInUser := s.Config.CurrentUsername

	if users == nil || len(users) == 0 {
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

