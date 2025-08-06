package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/cliffdoyle/social-network/internal/database"
)

type FollowRepository interface {
	InsertFollow(ctx context.Context, me string, foloweeID string) error
	DeleteFollow(ctx context.Context, me string, foloweeID string) error
	GetFollowers(ctx context.Context, me string) ([]string, error)
	GetPeopleIFollow(ctx context.Context, me string) ([]string, error)
}

type followRepository struct {
	DB *database.DB
}

func NewFollowingRepository(db *database.DB) FollowRepository {
	return &followRepository{DB: db}
}

func (f *followRepository) InsertFollow(ctx context.Context, me string, foloweeID string) error {
	query := `INSERT INTO  following VALUES (?,?)`
	_, err := f.DB.ExecContext(ctx, query, foloweeID, me)
	if err != nil {
		return fmt.Errorf("Failed to process follow")
	}
	return nil
}

func (f *followRepository) DeleteFollow(ctx context.Context, me string, foloweeID string) error {
	query := `DELETE from following f WHERE f.followee_id= ? AND f.follower_id=?`
	_, err := f.DB.ExecContext(ctx, query, foloweeID, me)
	if err != nil {
		return fmt.Errorf("Failed to process unfollow")
	}
	return nil
}

func (f *followRepository) GetFollowers(ctx context.Context, me string) ([]string, error) {
	query := `SELECT  follower_id FROM following WHERE followee_id =?`
	var peopleFollowingMe []string
	rows, err := f.DB.QueryContext(ctx, query, me)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("Querry deadline reached")
		}
		return nil, fmt.Errorf("error querrying database")
	}

	for rows.Next() {
		var userID string
		if err := rows.Scan(&userID); err != nil {
			return nil, fmt.Errorf("error querrying database")
		}
		peopleFollowingMe = append(peopleFollowingMe, userID)
	}
	return peopleFollowingMe, nil
}

func (f *followRepository) GetPeopleIFollow(ctx context.Context, me string) ([]string, error) {
	var peopleIFollow []string

	query := `SELECT foloweeID FROM following WHERE follower_id=?`

	rows, err := f.DB.QueryContext(ctx, query, me)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, fmt.Errorf("Querry deadline reached")
		}
		return nil, fmt.Errorf("error querrying database")
	}

	for rows.Next() {
		var userID string
		if err := rows.Scan(&userID); err != nil {
			return nil, fmt.Errorf("error querrying database")
		}
		peopleIFollow = append(peopleIFollow, userID)
	}
	return peopleIFollow, nil
}
