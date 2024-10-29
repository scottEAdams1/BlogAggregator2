package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scottEAdams1/BlogAggregator2/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arg) != 1 {
		return errors.New("wrong args")
	}

	url := cmd.arg[0]
	feed, err := s.db.GetFeedFromURL(context.Background(), url)
	if err != nil {
		return err
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return err
	}

	printFeedFollow(feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	if len(feedFollows) == 0 {
		fmt.Println("No feed follows")
		return nil
	}

	for _, feedFollow := range feedFollows {
		fmt.Println(feedFollow.FeedName)
	}
	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arg) != 1 {
		return errors.New("wrong args")
	}

	url := cmd.arg[0]

	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	})
	if err != nil {
		return err
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
