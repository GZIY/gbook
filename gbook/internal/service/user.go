package service

import (
	"context"
	"github.com/GZIY/gbook/gbook/internal/domain"
	"github.com/GZIY/gbook/gbook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	// 加密

	// 存储
	err := svc.repo.Create(ctx, u)
	if err != nil {
		return err
	}
	return nil
}
