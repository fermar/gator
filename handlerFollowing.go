package main

import (
	"context"
	"fmt"

	"github.com/fermar/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, usrInfo database.User) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("demasiados argumentos para el comando %v", cmd.name)
	}

	// usrInfo, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	// if err != nil {
	// 	return err
	// }

	feedsForUsr, err := s.db.GetFeedsFollowsForUser(context.Background(), usrInfo.ID)
	if err != nil {
		return err
	}
	fmt.Printf("Feeds del usuario %v:\n", s.conf.CurrentUserName)
	for _, feed := range feedsForUsr {
		fmt.Printf("\t* %v\n", feed.FeedName)
	}
	fmt.Println("-------------------")
	return nil
}
