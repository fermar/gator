package main

import (
	"context"
	"errors"
	"fmt"
)

// "context"
// "errors"
// "fmt"

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("demasiados argumentos para reset")
	}
	usuarios, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error al obtener usuarios: %w", err)
	}
	for _, usr := range usuarios {
		fmt.Printf("* %v", usr.Name)
		if usr.Name == s.conf.CurrentUserName {
			fmt.Print(" (current)")
		}
		fmt.Println()
	}
	// fmt.Printf("BD Reseteada\n")
	return nil
}
