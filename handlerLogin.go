package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New(fmt.Sprintf("Pocos argumentos para el comando %v", cmd.name))
	}
	// s.db.GetUser(context.Background(), id uuid.UUID)
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Printf("el usuario %v no existe en la base de datos\n", cmd.args[0])
		return err
	}

	err = s.conf.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("Usuario seteado...")
	return nil
}
