package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return &RSSFeed{}, err
	}

	req.Header.Set("User-Agent", "gator")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssfeed := RSSFeed{}
	err = xml.Unmarshal(data, &rssfeed)
	if err != nil {
		return &RSSFeed{}, err
	}

	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)
	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
	for i := range rssfeed.Channel.Item {
		rssfeed.Channel.Item[i].Description = html.UnescapeString(rssfeed.Channel.Item[i].Description)
		rssfeed.Channel.Item[i].Title = html.UnescapeString(rssfeed.Channel.Item[i].Title)
	}
	return &rssfeed, nil
}
