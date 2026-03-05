package main

import (
	"context"
	"fmt"

	"github.com/ScholarlyKiwi/gator/internal/database"
)

type handlerLoggedIn func(s *state, cmd command, user database.User) error
type handlerNormal func(s *state, cmd command) error

func middlewareLoggedIn(handler handlerLoggedIn) handlerNormal {

	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Error adding feed, unable to get user %v: %v\n", s.config.CurrentUserName, err)
		}
		return handler(s, cmd, user)
	}
}
