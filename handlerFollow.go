package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/fermar/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("pocos argumentos para el comando %v", cmd.name)
	}

	usrInfo, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}
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

	err = handlerLogin(s, cmd)
	if err != nil {
		return err
	}

	// err := s.conf.SetUser(cmd.args[0])
	// if err != nil {
	// 	return err
	// }
	fmt.Printf("Usuario registrado:\n,%v", usrData)
	return nil
}
