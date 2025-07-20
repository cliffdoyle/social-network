package service

import (
	"github.com/cliffdoyle/social-network/internal/models"
	"github.com/cliffdoyle/social-network/internal/repository"
)

type SessionService interface {
	PersistSession(id string) (string, error)
	GetSessionFromDB(sessionID string) (*models.Sessions,error)
}

type sessionService struct {
	sessionsRepo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) SessionService {
	return &sessionService{sessionsRepo: repo}
}

func (s *sessionService) PersistSession(id string) (string, error) {
	
}

func GetSessionFromDB(sessionID string) (*models.Sessions,error){

}