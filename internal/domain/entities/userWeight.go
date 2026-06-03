package entities

import "time"

type UserWeight struct {
	Id        int
	UserId    int
	Date      time.Time
	CreatedAt time.Time
	Value     float64
}

type UserWeightTelegram struct {
}

type UserWeightDb struct {
}
