package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("Error Agg command requires a request delay.")
	}

	time_between_args, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("Error invalid time as request delay, %v", cmd.arguments[0])
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_args)

	ticker := time.NewTicker(time_between_args)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
