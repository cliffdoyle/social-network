package service

import (
	"fmt"
	"time"

	"github.com/cliffdoyle/social-network/internal/models"
	"github.com/cliffdoyle/social-network/internal/repository"
	"github.com/google/uuid"
)

type SessionService interface {
	PersistSession(id string) (string, error)
	GetSessionFromDB(sessionID string) (*models.Sessions, error)
	DeleteSession(sessionID string)
}

type sessionService struct {
	sessionsRepo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) SessionService {
	return &sessionService{sessionsRepo: repo}
}

func (s *sessionService) PersistSession(id string) (string, error) {
	session := &models.Sessions{
		SessionID: uuid.New().String(),
		UserID:    id,
		Expires:   time.Now().Add(24 * time.Hour),
	}
	sessionID, err := s.sessionsRepo.InsertSession(session)
	if err != nil {
		return "", err
	}
	return sessionID, nil

}

func (s *sessionService) GetSessionFromDB(sessionID string) (*models.Sessions, error) {
	session, err := s.sessionsRepo.GetSessionID(sessionID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sessionService) DeleteSession(sessionID string) {
	err := s.sessionsRepo.DeleteSessionFromDB(sessionID)
	if err != nil {
		fmt.Println("failed to delete session.Please try again")
		return
	}
}
