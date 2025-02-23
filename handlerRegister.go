package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fermar/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New(fmt.Sprintf("Pocos argumentos para el comando %v", cmd.name))
	}
	var cur = database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}
	usrData, err := s.db.CreateUser(context.Background(), cur)
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
