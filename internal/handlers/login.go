package handlers

import (
	"context"
	"errors"
	"fmt"
)

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
