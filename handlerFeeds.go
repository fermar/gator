package main

import (
	"context"
	"errors"
	"fmt"
)

// "context"
// "errors"
// "fmt"

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("demasiados argumentos para reset")
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error al obtener usuarios: %w", err)
	}
	for _, feed := range feeds {
		fmt.Printf("%+v\n", feed)
	}
	return nil
}
