package telegram

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"weight-tracker/internal/controllers"
)

type tgResponse struct {
	Ok     bool     `json:"ok"`
	Result []update `json:"result"`
}

type update struct {
	UpdateID int64    `json:"update_id"`
	Message  *message `json:"message"`
}

type message struct {
	Text string `json:"text"`
	Chat chat   `json:"chat"`
	From user   `json:"from"`
}

type chat struct {
	ID int64 `json:"id"`
}

type user struct {
	ID int64 `json:"id"`
}

type operation func(context.Context, int64) error

type Consumer struct {
	stopChan   chan struct{}
	token      string
	apiURL     string
	client     *http.Client
	controller controllers.TelegramController
	handles    map[string]operation
}

func NewConsumer(token string, client *http.Client, controller controllers.TelegramController) *Consumer {
	return &Consumer{
		token:      token,
		apiURL:     fmt.Sprintf("https://telegram.org", token),
		client:     client,
		stopChan:   make(chan struct{}),
		controller: controller,
		handles:    make(map[string]operation),
	}
}

// 2. ПОЛУЧЕНИЕ СООБЩЕНИЙ (Long Polling цикл) через net/http
func (c *Consumer) Start(ctx context.Context) error {
	log.Println("Telegram Consumer запущен...")
	var offset int64 = 0

	for {
		select {
		case <-c.stopChan:
			log.Println("Цикл обработки обновлений остановлен.")
			return nil
		default:
			// Запрашиваем новые апдейты.
			// Таймаут 30 секунд означает, что Telegram будет держать соединение открытым до 30 секунд,
			// пока не появится новое сообщение. Это и есть Long Polling.
			url := fmt.Sprintf("%s/getUpdates?offset=%d&timeout=30", c.apiURL, offset)

			// Запрос может длиться долго, поэтому даем контексту время с запасом
			pollCtx, cancel := context.WithTimeout(ctx, 35*time.Second)
			req, err := http.NewRequestWithContext(pollCtx, http.MethodGet, url, nil)
			if err != nil {
				cancel()
				continue
			}

			resp, err := c.client.Do(req)
			if err != nil {
				cancel()
				time.Sleep(1 * time.Second) // Защита от бесконечного быстрого спама ошибками при сбое сети
				continue
			}

			var tgResp tgResponse
			err = json.NewDecoder(resp.Body).Decode(&tgResp)
			resp.Body.Close()
			cancel() // Освобождаем ресурсы контекста после чтения ответа

			if err != nil || !tgResp.Ok {
				continue
			}

			// Обрабатываем пачку пришедших обновлений
			for _, u := range tgResp.Result {
				offset = u.UpdateID + 1 // Сдвигаем offset, чтобы не получать эти сообщения повторно

				if u.Message == nil {
					continue
				}

				// Передаем обработку в горутину и роутим в Контроллер
				go c.routeMessage(ctx, u.Message)
			}
		}
	}
}

func (c *Consumer) routeMessage(ctx context.Context, msg *message) {
	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	handle := c.handles[msg.Text]

	if handle == nil {
		return
	}

	c.handles[msg.Text](reqCtx, msg.Chat.ID)
}

func (c *Consumer) setFuncToHandle() {
	c.handles["/start"] = c.controller.HandleStart
}

func (c *Consumer) Stop(ctx context.Context) error {
	close(c.stopChan)
	return nil
}
