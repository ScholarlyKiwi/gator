package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {

	dbuser, err := s.db.GetUsers(context.Background())

	if err != nil {
		fmt.Printf("Error listing users: %v\n", err)
	}
	for _, user := range dbuser {
		var current string
		if s.config.CurrentUserName == user.Name {
			current = "(current)"
		}
		fmt.Printf("* %v %v\n", user.Name, current)
	}

	return nil
}
