package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/fermar/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, usr database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("pocos argumentos para el comando %v", cmd.name)
	}
	// s.db.GetUser(context.Background(), id uuid.UUID)
	// usr, err := s.db.GetUser(context.Background(), s.conf.CurrentUserName)
	// if err != nil {
	// 	fmt.Printf("el usuario %v no existe en la base de datos\n", s.conf.CurrentUserName)
	// 	return err
	// }
	//
	// err = s.conf.SetUser(cmd.args[0])
	cfparams := database.CreateFeedsParams{}
	cfparams.ID = uuid.New()
	cfparams.CreatedAt = time.Now()
	cfparams.UpdatedAt = time.Now()
	cfparams.Name = cmd.args[0]
	cfparams.Url = cmd.args[1]
	cfparams.UserID = usr.ID
	// cfparams.UserID.UUID = usr.ID
	// cfparams.UserID.Valid = true

	feed, err := s.db.CreateFeeds(context.Background(), cfparams)
	if err != nil {
		return err
	}

	cmdFollow := command{name: "follow", args: []string{feed.Url}}

	err = handlerFollow(s, cmdFollow, usr)
	if err != nil {
		return err
	}

	fmt.Println("feed creado...")
	fmt.Printf("%+v\n", feed)
	return nil
}
