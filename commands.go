package main

import (
	"fmt"
	"os"
)

type commands struct {
	reg map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	call, ok := c.reg[cmd.command]
	if ok {
		if err := call(s, cmd); err != nil {
			return fmt.Errorf("Error running %v: %v\n", cmd.command, err)
		}
	} else {
		fmt.Printf("No such command %v\n", cmd.command)
		os.Exit(1)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) error {
	c.reg[name] = f
	return nil
}
