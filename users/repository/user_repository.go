package repository

import (
	"database/sql"

	_ "github.com/pastorilps/echo-swagger/app/docs"
	"github.com/pastorilps/echo-swagger/middleware"
	"github.com/pastorilps/echo-swagger/users/domain"
	"github.com/pastorilps/echo-swagger/users/entity"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepo(Conn *sql.DB) domain.UserRepository {
	return &userRepository{Conn}
}

func (u *userRepository) FetchUserBydId(id int16) (*entity.Users, error) {
	query := `select * from public.user  where id = $1`

	list, err := u.fetchUserById(query, id)
	if err != nil {
		return nil, err
	}

	pkg := &entity.Users{}
	if len(list) > 0 {
		pkg = list[0]
	} else {
		return nil, middleware.ErrorNotFound
	}

	return pkg, nil
}

func (u *userRepository) fetchUserById(query string, args ...interface{}) ([]*entity.Users, error) {
	rows, err := u.Conn.Query(query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()
	res := make([]*entity.Users, 0)
	for rows.Next() {
		str := new(entity.Users)
		err := rows.Scan(
			&str.ID,
			&str.Name,
			&str.Email,
			&str.Password,
			&str.Picture,
			&str.Newsletter,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		res = append(res, str)
	}

	return res, nil
}

func (u *userRepository) FetchAllUsers() ([]*entity.Users, error) {
	query := `select * from public.user order by id asc`

	return u.fetchAllUsers(query)
}

func (u *userRepository) fetchAllUsers(query string, args ...interface{}) ([]*entity.Users, error) {
	rows, err := u.Conn.Query(query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rows.Close()

	result := make([]*entity.Users, 0)
	for rows.Next() {
		us := new(entity.Users)
		err = rows.Scan(
			&us.ID,
			&us.Name,
			&us.Email,
			&us.Password,
			&us.Picture,
			&us.Newsletter,
		)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, us)
	}

	return result, nil
}
