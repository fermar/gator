package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("demasiados argumentos para el comando %v", cmd.name)
	}

	usrInfo, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}

	feedsForUsr, err := s.db.GetFeedsFollowsForUser(context.Background(), usrInfo.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Feeds del usuario %v:\n", s.conf.CurrentUserName)
	for _, feed := range feedsForUsr {
		fmt.Printf("\t* %v\n", feed.FeedName)
	}
	fmt.Println("-------------------")
	// if err != nil {
	// 	return err
	// }
	// cf := database.CreateFollowParams{
	// 	ID:        uuid.New(),
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// 	UserID:    usrInfo.ID,
	// 	FeedID:    feedInfo.ID,
	// }
	// followInfo, err := s.db.CreateFollow(context.Background(), cf)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("Usuario: %v\n", s.conf.CurrentUserName)
	// fmt.Printf("feed: %v\n", feedInfo.Name)
	// logging.Lg.Logger.Printf("follow creado:\n,%+v\n", followInfo)
	return nil
}
