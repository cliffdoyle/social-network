package repository

import "github.com/cliffdoyle/social-network/internal/database"

type SessionRepository interface{
	Insert()
	GetSessionID()
}

type sessionRepository struct{
	DB *database.DB
}

func NewSessionRepository( db *database.DB) SessionRepository{
	return &sessionRepository {DB:db}
}