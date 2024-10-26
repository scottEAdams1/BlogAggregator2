package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/scottEAdams1/BlogAggregator2/internal/config"
)

type state struct {
	cfgPointer *config.Config
}

type command struct {
	name string
	arg  []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	var state1 state
	state1.cfgPointer = &cfg

	cmds := commands{
		commandList: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 3 {
		log.Fatal("too few args provided")
	}

	cmd := command{
		name: args[1],
		arg:  args[2:],
	}

	err = cmds.run(&state1, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}

}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arg) == 0 {
		return errors.New("no username provided")
	}

	err := s.cfgPointer.SetUser(cmd.arg[0])
	if err != nil {
		return err
	}

	fmt.Println("Username set")
	return nil
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
