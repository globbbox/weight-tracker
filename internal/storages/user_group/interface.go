package user_group

import (
	"context"
	"weight-tracker/internal/domain/entities"
)

type UserGroupStorage interface {
	Save(ctx context.Context, userId int, groupId int) error
	BulkSave(ctx context.Context, groupId int, userIds []int) error
	GetByUserId(ctx context.Context, id int) ([]entities.Group, error)
	GetByGroupId(ctx context.Context, id int) ([]entities.User, error)
	DeleteById(ctx context.Context, id int) error
	DeleteByUserId(ctx context.Context, id int) error
	DeleteByGroupId(ctx context.Context, id int) error
}
