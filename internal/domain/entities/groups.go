package entities

type Group struct {
	Id         int
	TelegramId int64
	Name       string
	Users      []User
}

type GroupTelegram struct {
}

type GroupDb struct {
}
