package user_weight

import (
	"context"
	"weight-tracker/internal/domain/entities"
)

type UserWeightStorage interface {
	Save(ctx context.Context, value entities.UserWeight) error
	GetByUserId(ctx context.Context, userId int) ([]entities.UserWeight, error)
	GetById(ctx context.Context, id int) (entities.UserWeight, error)
	GetByDate(ctx context.Context, id int) (entities.UserWeight, error)
}
