package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) model.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (user *model.User, err error) {
	row := sq.Select("id", "name", "email", "password", "created_at").
		From("users").
		Where(sq.Eq{"email": email}).
		RunWith(r.db).
		QueryRowContext(ctx)
	var data model.User
	if err = row.Scan(
		&data.ID,
		&data.Name,
		&data.Email,
		&data.Password,
		&data.CreatedAt); err != nil {
		logrus.WithField("email", email).Error(err)
	}

	return &data, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (user *model.User, err error) {
	row := sq.Select("id", "name", "email", "role", "created_at").
		From("users").
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		QueryRowContext(ctx)
	var data model.User
	if err = row.Scan(
		&data.ID,
		&data.Name,
		&data.Email,
		&data.Role,
		&data.CreatedAt); err != nil {
		logrus.WithField("id", id).Error(err)
	}

	return &data, nil
}

func (r *userRepository) Create(ctx context.Context, data model.User) (newUser *model.User, err error) {
	timeNow := time.Now().UTC()

	result, err := sq.Insert("users").
		Columns("name", "email", "password", "created_at").
		Values(data.Name, data.Email, data.Password, timeNow).
		RunWith(r.db).
		ExecContext(ctx)
	if err != nil {
		logrus.WithField("data", data).Error(err)
		return
	}

	lastInsertID, _ := result.LastInsertId()

	newUser = &data
	newUser.ID = lastInsertID
	newUser.CreatedAt = timeNow

	return
}
