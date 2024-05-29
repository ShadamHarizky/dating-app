package repository

import (
	"github.com/ShadamHarizky/dating-app/domain/entity"
)

type UserRepository interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUser(uint64) (*entity.User, error)
	GetUsers([]int) ([]entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
	PurchasePremium(uint64) (*entity.User, error)
}
