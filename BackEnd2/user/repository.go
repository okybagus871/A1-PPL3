package user

import (
	"context"
	"database/sql"
	"github.com/go-kit/log"
	"errors"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type Repository interface {
	SignUp(ctx context.Context, user userProfile) error
	CheckEmail(ctx context.Context, email string) (bool, error)
	CheckUsername(ctx context.Context, username string) (bool, error)
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

func (r *repo) SignUp(ctx context.Context, user userProfile) error {
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

func (r *repo) CheckEmail(ctx context.Context, email string) (bool, error){
	var exists bool

	sql := `SELECT EXISTS(SELECT 1 FROM users WHERE email=$1);`

	if err := r.db.QueryRow(sql, email).Scan(&exists); err != nil {
		return false, nil
	}
	return exists, nil
}

func (r *repo) CheckUsername(ctx context.Context, username string) (bool, error){
	var exists bool

	sql := `SELECT EXISTS(SELECT 1 FROM users WHERE username=$1);`

	if err := r.db.QueryRow(sql, username).Scan(&exists); err != nil {
		return false, err
	}
	return exists, nil
}