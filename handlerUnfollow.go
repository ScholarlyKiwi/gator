package main

import (
	"context"
	"fmt"

	"github.com/ScholarlyKiwi/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Error: The register command requires a URL.")
	}

	url := cmd.arguments[0]
	feed, err := s.db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Error unfollowing %v: %v", url, err)
	}

	err_db := s.db.DeleteFeedFollows(context.Background(), database.DeleteFeedFollowsParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err_db != nil {
		return fmt.Errorf("Error in DB unfollowing %v: %v", url, err_db)
	}

	return nil
}
