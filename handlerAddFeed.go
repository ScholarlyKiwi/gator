package main

import (
	"context"
	"fmt"

	"github.com/ScholarlyKiwi/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Error addfeed requires two arguments, the name and url of the feed.")
	}

	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error adding feed, unable to get user %v: %v", s.config.CurrentUserName, err)
	}

	feedparams := database.AddFeedParams{
		Name: cmd.arguments[0],
		Url:  cmd.arguments[1],
		UserID: uuid.NullUUID{
			UUID:  user.ID,
			Valid: true},
	}

	feed, err := s.db.AddFeed(context.Background(), feedparams)
	if err != nil {
		return fmt.Errorf("Error adding feed: %v", err)
	}
	if feed.Name != cmd.arguments[0] {
		return fmt.Errorf("Error add feed name mismatch, \nentered: %v \nresult %v", feed.Name, cmd.arguments[0])
	}

	fmt.Println(feed)

	return nil
}
