package domain

import "github.com/pastorilps/echo-swagger/users/entity"

type UserUseCase interface {
	GetAllUsers() ([]*entity.Users, error)
}

type UserRepository interface {
	FetchAllUsers() ([]*entity.Users, error)
}
