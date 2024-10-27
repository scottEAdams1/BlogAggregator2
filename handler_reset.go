package main

import (
	"context"
	"log"
)

func handlerReset(s *state, cmd command) error {
	log.Println("Attempting to reset users...")
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		log.Println("Failed to reset users:", err)
		return err
	}
	log.Println("Users successfully reset.")
	return nil
}
