package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ScholarlyKiwi/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32
	if len(cmd.arguments) == 0 {
		limit = 2
	} else {
		value, err := strconv.ParseInt(cmd.arguments[0], 10, 32)
		if err != nil {
			return fmt.Errorf("Error browsing %v is not a number\n", cmd.arguments[0])
		}
		limit = int32(value)
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Limit:  limit,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("Error getting posts: %v", err)
	}
	if len(posts) == 0 {
		fmt.Println("No posts to browse.")
	} else {
		for _, post := range posts {
			fmt.Println("-----------------------------")
			fmt.Printf("Title: %v\n", post.Title)
			fmt.Printf("URL: %v\n", post.Url)
			fmt.Printf("Description: %v\n", post.Description)
			fmt.Printf("Updated at: %v\n", post.UpdatedAt)
		}
	}
	return nil
}
