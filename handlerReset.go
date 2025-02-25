package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("demasiados argumentos para reset")
	}
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error en el reset: %w", err)
	}
	fmt.Printf("BD Reseteada\n")
	return nil
}
