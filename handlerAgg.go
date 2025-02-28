package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// "context"
// "errors"
// "fmt"

func handlerAdd(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("demasiados argumentos para agg")
	}

	// usuarios, err := s.db.GetUsers(context.Background())
	// if err != nil {
	// 	return fmt.Errorf("error al obtener usuarios: %w", err)
	// }
	// for _, usr := range usuarios {
	// 	fmt.Printf("* %v", usr.Name)
	// 	if usr.Name == s.conf.CurrentUserName {
	// 		fmt.Print(" (current)")
	// 	}
	// 	fmt.Println()
	// }
	// // fmt.Printf("BD Reseteada\n")
	return nil
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
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
