package telegram

import (
	"context"
	"weight-tracker/internal/controllers"
)

type Consumer struct {
	stopChan   chan struct{}
	controller controllers.TelegramController
}

func (n Consumer) Start(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

//func (c *Consumer) routeMessage(ctx context.Context, msg *tgbotapi.Message) {
//	reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
//	defer cancel()
//
//	// Проверяем, что ввел пользователь, и вызываем нужный метод Контроллера
//	switch msg.Text {
//	case "/start":
//		c.controller.HandleStart(reqCtx, msg)
//	default:
//		// Если это не команда, считаем, что пользователь вводит вес
//		c.controller.HandleSaveWeight(reqCtx, msg)
//	}
//}

func (n Consumer) Stop(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
