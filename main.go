package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/scottEAdams1/BlogAggregator2/internal/config"
	"github.com/scottEAdams1/BlogAggregator2/internal/database"
)

type state struct {
	cfgPointer *config.Config
	db         *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		cfgPointer: &cfg,
		db:         dbQueries,
	}

	cmds := commands{
		commandList: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	args := os.Args
	if len(args) < 2 {
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
