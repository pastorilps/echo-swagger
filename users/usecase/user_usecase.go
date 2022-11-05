package usecase

import (
	"github.com/pastorilps/echo-swagger/users/domain"
	"github.com/pastorilps/echo-swagger/users/entity"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(ur domain.UserRepository) domain.UserUseCase {
	return &userUseCase{
		userRepo: ur,
	}
}

func (u *userUseCase) GetAllUsers() ([]*entity.Users, error) {
	list, err := u.userRepo.FetchAllUsers()
	if err != nil {
		return nil, err
	}

	return list, nil
}
