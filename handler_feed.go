package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scottEAdams1/BlogAggregator2/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	username := s.cfgPointer.CurrentUserName
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return err
	}

	if len(cmd.arg) != 2 {
		return errors.New("missing args")
	}

	feedName := cmd.arg[0]
	url := cmd.arg[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}

func handlerFeed(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("No feeds")
		return nil
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserByID(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(user.Name)
	}

	return nil
}
