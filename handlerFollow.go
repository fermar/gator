package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/fermar/gator/internal/database"
	"github.com/fermar/gator/internal/logging"
)

func handlerFollow(s *state, cmd command, usrInfo database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("pocos argumentos para el comando %v", cmd.name)
	}
	// usrInfo, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	// if err != nil {
	// 	return err
	// }
	feedInfo, err := s.db.GetFeedsByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	cf := database.CreateFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    usrInfo.ID,
		FeedID:    feedInfo.ID,
	}
	followInfo, err := s.db.CreateFollow(context.Background(), cf)
	if err != nil {
		return err
	}
	fmt.Printf("Usuario: %v\n", s.conf.CurrentUserName)
	fmt.Printf("feed: %v\n", feedInfo.Name)
	logging.Lg.Logger.Printf("follow creado:\n,%+v\n", followInfo)
	return nil
}
