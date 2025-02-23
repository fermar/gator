package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New(fmt.Sprintf("Pocos argumentos para el comando %v", cmd.name))
	}
	err := s.conf.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("Usuario seteado...")
	return nil
}
