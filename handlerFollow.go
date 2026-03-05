package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ScholarlyKiwi/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("Error: follow command requires a URL\n")
	}
	url := cmd.arguments[0]

	feed, err_feed := s.db.GetFeedByURL(context.Background(), url)
	if err_feed != nil {
		return fmt.Errorf("Error following feed %v: %v\n", url, err_feed)
	}

	user, err_user := s.db.GetUser(context.Background(), user.Name)
	if err_user != nil {
		return fmt.Errorf("Error following feed user error: %v\n", err_user)
	}

	feed_follows, err_ff := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID},
	)
	if err_ff != nil {
		return fmt.Errorf("Error creating feed follows: %v", err_ff)
	}
	fmt.Printf("Feed Follows:\n- User: %v\n- Feed: %v\n", feed_follows.Username, feed_follows.Feedname)

	return nil
}
