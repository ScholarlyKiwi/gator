package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ScholarlyKiwi/gator/internal/config"
	"github.com/ScholarlyKiwi/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {

	fmt.Println("Welcome to gator")

	var s state
	var c commands

	if err := configuration(&c, &s); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Error - no command supplied")
		os.Exit(1)
	}

	args := os.Args

	cmd := command{
		command: args[1],
	}
	if len(args) > 1 {
		cmd.arguments = args[2:]
	}

	if err := c.run(&s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func configuration(c *commands, s *state) error {

	gatorConfig, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	s.config = &gatorConfig

	c.reg = make(map[string]func(*state, command) error)
	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerUsers)

	db, err := sql.Open("postgres", s.config.DBurl)
	s.db = database.New(db)

	return nil
}
