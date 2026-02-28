package main

import (
	"github.com/ScholarlyKiwi/gator/internal/config"
	"github.com/ScholarlyKiwi/gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
