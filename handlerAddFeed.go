package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ScholarlyKiwi/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 2 {
		return fmt.Errorf("Error addfeed requires two arguments, the name and url of the feed.\n")
	}

	feedparams := database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
		Url:       cmd.arguments[1],
		UserID:    user.ID,
	}

	feed, err := s.db.AddFeed(context.Background(), feedparams)
	if err != nil {
		return fmt.Errorf("Error adding feed: %v\n", err)
	}
	if feed.Name != cmd.arguments[0] {
		return fmt.Errorf("Error adding feed name mismatch, \nentered: %v \nresult %v", feed.Name, cmd.arguments[0])
	}

	fmt.Printf("Added Feed %v for user %v\n", feed.Name, s.config.CurrentUserName)

	followcmd := command{
		command:   "follow",
		arguments: []string{cmd.arguments[1]},
	}
	err = handlerFollow(s, followcmd, user)
	if err != nil {
		return err
	}

	return nil
}
