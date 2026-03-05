package main

import (
	"context"
	"fmt"

	"github.com/ScholarlyKiwi/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) > 0 {
		return fmt.Errorf("Error: following command does not requires arguments\n")
	}

	feed_follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("Error retireving feeds for user %v: %v\n", user.Name, err)
	}

	fmt.Printf("Feeds for user %v\n", user.Name)
	for _, feed := range feed_follows {
		fmt.Printf("- %v\n", feed.Feedname)
	}

	return nil
}
