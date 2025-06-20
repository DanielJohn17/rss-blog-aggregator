package handlers

import (
	"errors"
	"fmt"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return errors.New("The login handler expects a single argument")
	}

	s.Config.SetUser(cmd.Args[0])

	fmt.Printf("User %s has been set.\n", cmd.Args[0])

	return nil
}
