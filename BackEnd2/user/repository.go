package user

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	"errors"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type Repository interface {
	CreateUser(ctx context.Context, user userProfile) error
}

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "postgres"),
	}
}

func (r *repo) CreateUser(ctx context.Context, user userProfile) error {
	sql := `INSERT INTO users (username, name, password, created_date, email)
			VALUES ($1, $2, $3, $4, $5)`

	if user.Name == "" || user.Password == "" || user.Username == "" || user.Email == ""{
		return RepoErr
	}
	_, err := r.db.ExecContext(
		ctx, sql, user.Username, user.Name, user.Password, user.Created_date, user.Email)
	
	if err != nil {
		return err
	}

	return nil
}
