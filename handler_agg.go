package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/scottEAdams1/BlogAggregator2/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arg) != 1 {
		return errors.New("wrong args")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arg[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s...", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
		if err != nil {
			log.Printf("error scraping feeds: %v", err)
			continue
		}
	}
}

func scrapeFeeds(s *state) error {
	fmt.Println("Starting feed scrape...")
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("1")
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Println("2")
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Println("3")
		return err
	}
	fmt.Printf("Found %d items in feed", len(rssFeed.Channel.Item))

	for _, item := range rssFeed.Channel.Item {
		fmt.Printf("Attempting to save post: %s", item.Title)
		publishedAt := sql.NullTime{}
		if t, err := time.Parse(time.RFC1123Z, item.PubDate); err == nil {
			publishedAt = sql.NullTime{
				Time:  t,
				Valid: true,
			}
		}

		_, err = s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: publishedAt,
			FeedID:      feed.ID,
		})
		if err != nil {
			fmt.Printf("failed to create post: %v", err)
			continue
		}
		fmt.Printf("Successfully saved post: %s", item.Title)

	}
	return nil
}
