package repository

import (
	"database/sql"

	_ "github.com/pastorilps/echo-swagger/app/docs"
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

func (u *userRepository) FetchAllUsers() ([]*entity.Users, error) {
	query := `select id, name, email, password, picture, newsletter from public.user order by id asc`

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
