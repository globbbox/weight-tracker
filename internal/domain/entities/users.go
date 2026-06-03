package entities

type User struct {
	Id         int
	ExternalID string
	Source     string
	Name       string
}

type UserTelegram struct {
	//данные получаемые от телеграмма
}

type UserDb struct {
	Id         int
	TelegramId int64
	Name       string
}
