package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	comandos map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) error {

	c.comandos[name] = f
	return nil
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.comandos[cmd.name]
	if !ok {
		return fmt.Errorf("El comando %v no existe", cmd.name)
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}
