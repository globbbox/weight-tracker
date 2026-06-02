package interfaces

import (
	"weight-tracker/internal/domain/entities"
)

type Users interface {
	Save(user entities.User) error
	GetById(id int) (entities.UserDb, error)
	GetByTgId(id int64) (entities.UserDb, error)
	UpdateById(id int, user entities.User) error
	UpdateByTgId(id int64, user entities.User) error
	Delete(user entities.User) error
}
