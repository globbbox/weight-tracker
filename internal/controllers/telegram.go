package controllers

import (
	"weight-tracker/internal/domain/entities"
	"weight-tracker/internal/services"
)

type TelegramController struct {
	UsersService services.UsersService
}

func NewTelegramController(usersStorage services.UsersService) *TelegramController {
	return &TelegramController{UsersService: usersStorage}
}

func (t *TelegramController) NewUser(user entities.User) {

}
