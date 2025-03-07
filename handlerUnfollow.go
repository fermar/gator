package main

import (
	"context"
	"fmt"

	"github.com/fermar/gator/internal/database"
	"github.com/fermar/gator/internal/logging"
)

func handlerUnfollow(s *state, cmd command, usrInfo database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("pocos argumentos para el comando %v", cmd.name)
	}

	feedInfo, err := s.db.GetFeedsByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return err
	}
	df := database.DeleteFollowParams{
		UserID: usrInfo.ID,
		FeedID: feedInfo.ID,
	}
	err = s.db.DeleteFollow(context.Background(), df)
	if err != nil {
		return err
	}
	fmt.Println("Unfollow para:")
	fmt.Printf("Usuario: %v\n", s.conf.CurrentUserName)
	fmt.Printf("feed: %v\n", feedInfo.Name)
	logging.Lg.Logger.Printf("unfollow Ejecutado\n")
	return nil
}
