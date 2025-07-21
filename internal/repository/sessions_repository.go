package repository

import (
	"fmt"

	"github.com/cliffdoyle/social-network/internal/database"
	"github.com/cliffdoyle/social-network/internal/models"
)

type SessionRepository interface {
	InsertSession(session *models.Sessions) (string, error)
	GetSessionID(id string) (*models.Sessions, error)
	DeleteSessionFromDB(sessionID string) error
}

type sessionRepository struct {
	DB *database.DB
}

func NewSessionRepository(db *database.DB) SessionRepository {
	return &sessionRepository{DB: db}
}

func (r *sessionRepository) InsertSession(session *models.Sessions) (string, error) {
	query := `
        INSERT INTO sessions (sessionID,userID,Expires)
        VALUES (?, ?, ?)
		`

	// Execute the query and get the generated ID
	_, err := r.DB.Exec(query,
		session.SessionID,
		session.UserID,
		session.Expires,
	)
	if err != nil {
		return "", fmt.Errorf("failed to insert to database")
	}
	return session.SessionID, nil

}

func (r *sessionRepository) GetSessionID(id string) (*models.Sessions, error) {
	query := `SELECT sessionID, userID, expires FROM Sessions WHERE id=?`
	row := r.DB.QueryRow(query, id)
	var sess *models.Sessions
	err := row.Scan(sess.SessionID, &sess.UserID, &sess.Expires)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (r *sessionRepository) DeleteSessionFromDB(sessionID string) error {
	query := "DELETE FROM Sessions WHERE id=?"
	_, err := r.DB.Exec(query, sessionID)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}
