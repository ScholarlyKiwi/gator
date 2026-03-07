package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/ScholarlyKiwi/gator/internal/database"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
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

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	var feed *RSSFeed
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	request, err_req := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	if err_req != nil {
		return feed, fmt.Errorf("Error creating new quest: %v", err_req)
	}

	request.Header.Set("User-Agent", "gator")
	res, err_res := client.Do(request)
	if err_res != nil {
		return feed, fmt.Errorf("Error doing request: %v", err_res)
	}
	defer res.Body.Close()

	body, err_body := io.ReadAll(res.Body)
	if err_body != nil {
		return feed, fmt.Errorf("Error reading body: %v", err_res)
	}

	err_xml := xml.Unmarshal(body, &feed)
	if err_xml != nil {
		return feed, fmt.Errorf("Error in Feed XML: %v", err_xml)
	}

	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	for idx := range feed.Channel.Item {
		feed.Channel.Item[idx].Title = html.UnescapeString(feed.Channel.Item[idx].Title)
		feed.Channel.Item[idx].Description = html.UnescapeString(feed.Channel.Item[idx].Description)
	}

	return feed, nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching feeds: %v", err)
	}
	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true}})

	rssfeed, err := fetchFeed(context.Background(), feed.Url)

	for _, feedItems := range rssfeed.Channel.Item {

		parsedDate, err := dateparse.ParseStrict(feedItems.PubDate)
		if err != nil {
			fmt.Printf("Unable to parse date %v on post %v: %v\n", feedItems.PubDate, feedItems.Title, err)
		} else {
			postParam := database.CreatePostParams{
				ID:          uuid.New(),
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
				Title:       feedItems.Title,
				Url:         feedItems.Link,
				PublishedAt: parsedDate,
				FeedID:      feed.ID,
			}
			_, err := s.db.CreatePost(context.Background(), postParam)
			if err != nil {
				fmt.Printf("Error creating post %v: %v", feedItems.Title, err)
			}
		}
	}

	return nil
}
