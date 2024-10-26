package main

import (
	"errors"
	"fmt"
)

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
