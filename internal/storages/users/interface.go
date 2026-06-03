package users

import (
	"context"
	"weight-tracker/internal/domain/entities"
)

type UsersStorage interface {
	Save(ctx context.Context, user entities.User) error
	GetById(ctx context.Context, id int) (entities.User, error)
	GetByTgId(ctx context.Context, id int64) (entities.User, error)
	UpdateById(ctx context.Context, id int, user entities.User) error
	UpdateByTgId(ctx context.Context, id int64, user entities.User) error
	DeleteById(ctx context.Context, id int, user entities.User) error
	DeleteByTgId(ctx context.Context, id int64, user entities.User) error
}
