package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/scottEAdams1/BlogAggregator2/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.arg) > 1 {
		return errors.New("wrong args")
	}

	limit := 2
	if len(cmd.arg) == 1 {
		specifiedLimit, err := strconv.Atoi(cmd.arg[0])
		limit = specifiedLimit
		if err != nil {
			return err
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%s from \n", post.PublishedAt.Time.Format("Mon Jan 2"))
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}

	return nil

}
