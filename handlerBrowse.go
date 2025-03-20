package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fermar/gator/internal/database"
)

// "context"
// "errors"
// "fmt"

func handlerBrowse(s *state, cmd command, usrInfo database.User) error {
	var limite int32 = 2
	if len(cmd.args) > 0 {
		lim, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return err
		}
		limite = int32(lim)
		// return errors.New("demasiados argumentos para browse")
	}

	usr, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	if err != nil {
		return err
	}
	// feeds, err := s.db.GetFeedsFollowsForUser(context.Background(), usr.ID)
	// if err != nil {
	// 	return err
	// }
	gpfu := database.GetPostsFromUserParams{UserID: usr.ID, Limit: limite}

	posts, err := s.db.GetPostsFromUser(context.Background(), gpfu)
	if err != nil {
		return err
	}

	for _, post := range posts {
		fmt.Printf("%+v\n", post.Title)
	}
	return nil
}
