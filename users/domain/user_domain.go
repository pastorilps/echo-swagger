package domain

import (
	"github.com/pastorilps/echo-swagger/users/entity"
)

type UserUseCase interface {
	GetAllUsers() ([]*entity.Users, error)
	GetUserById(id int16) (*entity.Users, error)
}

type UserRepository interface {
	FetchAllUsers() ([]*entity.Users, error)
	FetchUserBydId(id int16) (*entity.Users, error)
}
