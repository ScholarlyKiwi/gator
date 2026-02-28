package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ScholarlyKiwi/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Error: The register command requires exactly one argument.")
	}

	name := cmd.arguments[0]

	dbuser, err := s.db.GetUser(context.Background(), cmd.arguments[0])

	if dbuser.Name == name {
		fmt.Println("Unable to create user, user already exists")
		os.Exit(1)
	}

	_, err = s.db.CreateUser(context.Background(),
		database.CreateUserParams{ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      name})
	if err != nil {
		return fmt.Errorf("Error creating user: %v", err)
	}
	s.config.SetUser(name)
	fmt.Printf("User %v was created\n", name)

	return nil
}
