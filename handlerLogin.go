package main

import (
	"context"
	"fmt"
	"os"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("Error: login command requires one argument")
	}

	dbuser, _ := s.db.GetUser(context.Background(), cmd.arguments[0])
	if dbuser.Name != cmd.arguments[0] {
		fmt.Printf("Unable to login, user %v does not exists\n", cmd.arguments[0])
		os.Exit(1)
	}

	s.config.SetUser(cmd.arguments[0])
	fmt.Printf("User has been set to %v\n", s.config.CurrentUserName)
	return nil
}
