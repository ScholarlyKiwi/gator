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

		if err_feed.Error() == "sql: no rows in result set" {
			handlerAddFeed(s,
				command{
					command:   "addfeed",
					arguments: []string{url, url},
				},
				user)
			feed, err_feed = s.db.GetFeedByURL(context.Background(), url)
		}
		if err_feed != nil {
			return fmt.Errorf("Error following feed %v: %v\n", url, err_feed)
		}
	}

	user, err_user := s.db.GetUser(context.Background(), user.Name)
	if err_user != nil {
		return fmt.Errorf("Error following feed user error: %v\n", err_user)
	}

	userfeed_follow, err_userfeed := s.db.GetFeedFollowsForUserFeed(context.Background(),
		database.GetFeedFollowsForUserFeedParams{
			Name: user.Name,
			Url:  url,
		})

	found_feed := false
	if err_userfeed != nil {
		if err_userfeed.Error() != "sql: no rows in result set" {
			return fmt.Errorf("Error checking for existing feed %v for user %v: %v\n", url, user.Name, err_userfeed)
		}
	}
	if userfeed_follow.FeedID != feed.ID {
		return fmt.Errorf("Error checking for exiting feed mismatch, \nexpected: %v\nresult: %v\n", feed.Name, userfeed_follow.Feedname)
	} else {
		found_feed = true
	}

	if found_feed {
		fmt.Printf("User %v is already following %v\n", userfeed_follow.Username, userfeed_follow.Feedname)
	} else {
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
	}

	return nil
}
