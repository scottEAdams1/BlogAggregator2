package main

import (
	"context"

	"github.com/scottEAdams1/BlogAggregator2/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		username := s.cfgPointer.CurrentUserName
		user, err := s.db.GetUser(context.Background(), username)
		if err != nil {
			return err
		}

		return handler(s, cmd, user)
	}
}
