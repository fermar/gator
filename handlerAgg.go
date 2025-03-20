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

	"github.com/google/uuid"
	"github.com/lib/pq"

	"github.com/fermar/gator/internal/database"
	"github.com/fermar/gator/internal/logging"
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
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
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
	logging.Lg.Logger.Printf("fetchfeed para %v", nextFeed.Url)
	rssFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return err
	}

	logging.Lg.Logger.Printf("feeds encontrados: %v", len(rssFeed.Channel.Item))
	mffparams := database.MarkFeedFetchedParams{
		UpdatedAt:     time.Now(),
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:            nextFeed.ID,
	}
	err = s.db.MarkFeedFetched(context.Background(), mffparams)
	if err != nil {
		return err
	}
	fmt.Printf("Buscando posts para %v:\n", rssFeed.Channel.Title)

	postParams := database.CreatePostParams{}
	dups := 0
	for _, feedItem := range rssFeed.Channel.Item {
		postParams.ID = uuid.New()
		postParams.CreatedAt = time.Now()
		postParams.UpdatedAt = time.Now()
		postParams.Title = feedItem.Title
		postParams.Url = feedItem.Link
		postParams.Description = sql.NullString{String: feedItem.Description, Valid: true}
		// pubDat, err := time.Parse("Mon, 02 Jan 2006 15:04:05 -0700", feedItem.PubDate)
		pubDat, err := time.Parse(time.RFC1123Z, feedItem.PubDate)
		if err != nil {
			return err
		}
		postParams.PublishedAt = pubDat
		postParams.FeedID = nextFeed.ID
		// fmt.Println("-------feeitem-----")
		// fmt.Printf("%+v", feedItem)
		post, err := s.db.CreatePost(context.Background(), postParams)
		if err != nil {
			var pqerr *pq.Error
			if errors.As(err, &pqerr) {
				pqerr = err.(*pq.Error)
				if pqerr.Code != "23505" {
					return err
				} else {
					dups++
					continue
				}
			}
		}
		fmt.Printf("Post creado en bd, titulo: %v\n", post.Title)
		// fmt.Printf("\t - %v\n", feedItem.Title)
	}
	fmt.Printf("Post duplicados: %v\n", dups)
	fmt.Println("------------")
	return nil
}
