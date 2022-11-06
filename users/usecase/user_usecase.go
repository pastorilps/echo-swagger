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

func (u *userUseCase) GetUserById(id int16) (*entity.Users, error) {
	res, err := u.userRepo.FetchUserBydId(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *userUseCase) GetAllUsers() ([]*entity.Users, error) {
	list, err := u.userRepo.FetchAllUsers()
	if err != nil {
		return nil, err
	}

	return list, nil
}
