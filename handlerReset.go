package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {

	if err := s.db.DeleteUsers(context.Background()); err != nil {
		fmt.Printf("Error reseting users: %v", err)
		os.Exit(1)
	}
	fmt.Println("Successfully reset users")
	return nil
}
