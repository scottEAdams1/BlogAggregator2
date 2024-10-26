package main

import (
	"log"
	"os"

	"github.com/scottEAdams1/BlogAggregator2/internal/config"
)

type state struct {
	cfgPointer *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	programState := &state{
		cfgPointer: &cfg,
	}

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

	err = cmds.run(programState, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}

}
