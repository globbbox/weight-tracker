package services

import (
	"context"
	"weight-tracker/internal/domain/entities"
)

type NotificationSender interface {
	SendMessage(ctx context.Context, user entities.User, text string)
}
