package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	arg  []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandList[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if s.cfgPointer == nil {
		return errors.New("config not loaded")
	}

	if c.commandList[cmd.name] == nil {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}

	return c.commandList[cmd.name](s, cmd)
}
