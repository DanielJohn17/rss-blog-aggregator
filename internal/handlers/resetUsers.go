package handlers

import (
	"context"
	"fmt"
)

func ResetUsersHandler(s *State, _ Command) error {
	cxt := context.Background()

	if err := s.Db.Deleteusers(cxt); err != nil {
		return fmt.Errorf("Error deleting users")
	}

	fmt.Println("All users have been reset.")
	return nil
}
