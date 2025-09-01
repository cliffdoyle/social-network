package main

import (
	"database/sql"

	"github.com/cliffdoyle/social-network/internal/repository"
	"github.com/cliffdoyle/social-network/internal/service"
)

func newReactionHandler(db *sql.DB) *ReactionHandler {
	repo := repository.NewReactionRepository(db)
	service := service.NewReactionService(repo)
	return NewReactionHandler(service)
}
