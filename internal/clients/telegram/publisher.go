package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

func (c *Consumer) SendTextMessage(chatID int64, text string) error {
	url := fmt.Sprintf("%s/sendMessage", c.apiURL)

	reqData := sendMessageReq{
		ChatID: chatID,
		Text:   text,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return err
	}

	// Делаем POST-запрос с JSON. Используем Context для безопасности сетевого запроса
	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("telegram API returned status: %s", resp.Status)
	}

	return nil
}
