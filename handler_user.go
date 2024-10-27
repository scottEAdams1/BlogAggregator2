package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/scottEAdams1/BlogAggregator2/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arg) != 1 {
		return errors.New("no username provided")
	}

	_, err := s.db.GetUser(context.Background(), cmd.arg[0])
	if err != nil {
		return err
	}

	err = s.cfgPointer.SetUser(cmd.arg[0])
	if err != nil {
		return err
	}

	fmt.Println("Username set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arg) != 1 {
		return errors.New("no name given")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.arg[0],
	})
	if err != nil {
		return err
	}

	err = s.cfgPointer.SetUser(user.Name)
	if err != nil {
		return err
	}
	fmt.Println("User created")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return nil
	}

	for i := 0; i < len(users); i++ {
		if s.cfgPointer.CurrentUserName == users[i].Name {
			fmt.Println("* " + users[i].Name + " (current)")
		} else {
			fmt.Println("* " + users[i].Name)
		}
	}
	return nil
}
