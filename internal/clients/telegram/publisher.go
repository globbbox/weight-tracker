package telegram

import (
	"context"
	"net/http"
	"time"
	"weight-tracker/internal/domain/entities"
)

type Publisher struct {
	token  string
	chatID int64
	client *http.Client
}

func NewPublisher(botToken string, chatID int64) *Publisher {
	return &Publisher{
		token:  botToken,
		chatID: chatID,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
	}
}

func (t Publisher) SendMessage(ctx context.Context, user entities.User, text string) {
	//TODO implement me
	panic("implement me")
}
