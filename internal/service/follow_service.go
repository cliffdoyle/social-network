package service

import (
	"context"

	"github.com/cliffdoyle/social-network/internal/repository"
)

type FollowService interface {
	FollowUser(ctx context.Context, me string, foloweeID string) error
	UnFollowUser(ctx context.Context, me string, foloweeID string) error
	ShowFollowers(ctx context.Context, me string) ([]string, error)
	ShowPeopleIFollow(ctx context.Context, me string) ([]string, error)
}

type followService struct {
	followRepo repository.FollowRepository
}

func NewFollowService(repo repository.FollowRepository) FollowService {
	return &followService{followRepo: repo}
}

func (f *followService) FollowUser(ctx context.Context, me string, foloweeID string) error {
	err := f.followRepo.InsertFollow(ctx, foloweeID, me)
	if err != nil {
		return err
	}
	return nil
}

func (f *followService) UnFollowUser(ctx context.Context, me string, foloweeID string) error {
	err := f.followRepo.DeleteFollow(ctx, foloweeID, me)
	if err != nil {
		return err
	}
	return nil
}

func (f *followService) ShowFollowers(ctx context.Context, me string) ([]string, error) {
	followersIDS, err := f.followRepo.GetFollowers(ctx, me)
	if err != nil {
		return nil, err
	}
	return followersIDS, nil
}

func (f *followService) ShowPeopleIFollow(ctx context.Context, me string) ([]string, error) {
	followeeIDS, err := f.followRepo.GetPeopleIFollow(ctx, me)
	if err != nil {
		return nil, err
	}
	return followeeIDS, nil
}
