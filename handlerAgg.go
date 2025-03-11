package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/fermar/gator/internal/database"
)

// "context"
// "errors"
// "fmt"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("falta parámetro")
	}

	tbreqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Collecting feeds every %v\n", tbreqs)
	ticker := time.NewTicker(tbreqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
	// url := "https://www.wagslane.dev/index.xml"
	// rssFeed, err := fetchFeed(context.Background(), url)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v\n", rssFeed)
	//
	// return nil
}

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
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	client := http.Client{Timeout: 10 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	xmlBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	retFeed := RSSFeed{}
	xml.Unmarshal(xmlBody, &retFeed)
	retFeed.Channel.Title = html.UnescapeString(retFeed.Channel.Title)
	retFeed.Channel.Description = html.UnescapeString(retFeed.Channel.Description)
	for i, itemFeed := range retFeed.Channel.Item {
		itemFeed.Title = html.UnescapeString(itemFeed.Title)
		itemFeed.Description = html.UnescapeString(itemFeed.Description)
		retFeed.Channel.Item[i] = itemFeed
	}
	return &retFeed, nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}
	mffparams := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            nextFeed.ID,
	}
	err = s.db.MarkFeedFetched(context.Background(), mffparams)
	if err != nil {
		return err
	}
	fmt.Println("Titulos:")
	for _, feedItem := range rssFeed.Channel.Item {
		fmt.Printf("\t - %v\n", feedItem.Title)
	}
	return nil
}
