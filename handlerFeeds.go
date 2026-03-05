package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error retrieving feeds: %v", err)
	}

	fmt.Println("Listing all Feeds:")

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Error retrieving user for feeds: %v", err)
		}
		fmt.Printf("**%v**\n", feed.Name)
		fmt.Printf("- url: %v\n", feed.Url)
		fmt.Printf("- username: %v\n", user.Name)

	}

	return nil
}
