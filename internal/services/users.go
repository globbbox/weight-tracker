package services

import (
	"context"
	"weight-tracker/internal/domain/entities"
	"weight-tracker/internal/storages/users"
)

type UsersService struct {
	UsersStorage users.UsersStorage
}

func NewUsersService(usersStorage users.UsersStorage) *UsersService {
	return &UsersService{UsersStorage: usersStorage}
}

func (s *UsersService) CreateUserProfile(ctx context.Context, user entities.User) error {

	err := s.UsersStorage.Save(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
