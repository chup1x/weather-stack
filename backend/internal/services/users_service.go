package services

import (
	"context"
	"fmt"

	"github.com/chup1x/weather-stack/internal/domain"
	"github.com/google/uuid"
)

type userCreater interface {
	Create(context.Context, *domain.UserEntity) error
}

type userProvider interface {
	SelectByID(context.Context, domain.UserID) (*domain.UserEntity, error)
	SelectByTelegramID(context.Context, int64) (*domain.UserEntity, error)
}

type userStorage interface {
	userCreater
	userProvider
}

type UserService struct {
	users userStorage
}

func NewUserService(repo userStorage) *UserService {
	return &UserService{users: repo}
}

func (s *UserService) CreateUser(ctx context.Context, new *domain.UserEntity) (domain.UserID, error) {
	newID := uuid.New()
	new.ID = domain.UserID{UUID: newID}
	if err := s.users.Create(ctx, new); err != nil {
		return domain.UserID{}, fmt.Errorf("to create a user: %w", err)
	}
	return new.ID, nil
}

func (s *UserService) GetProfileByTelegramID(ctx context.Context, id int64) (*domain.UserEntity, error) {
	user, err := s.users.SelectByTelegramID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("to select a user by telegram id: %w", err)
	}
	return user, nil
}

func (s *UserService) GetProfileByID(ctx context.Context, id domain.UserID) (*domain.UserEntity, error) {
	user, err := s.users.SelectByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("to select a user by id: %w", err)
	}
	return user, nil
}
