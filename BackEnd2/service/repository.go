package service

import (
	// "strconv"
	// "math/rand"
	// "time"
	"context"
	"database/sql"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"errors"
	"BackEnd/datastruct"
)

var RepoErr = errors.New("Unable to handle Repo Request")

type Repository interface {
	SignUp(ctx context.Context, user datastruct.User) error
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

func (r *repo) SignUp(ctx context.Context, user datastruct.User) error {
	// token_hash := GenerateRandomString(15)
	// user.Token_hash = token_hash
	
	// min := 1000
	// max := 9999
	// rand.Seed(time.Now().UnixNano())
	// rd := rand.Intn(max-min+1) + min
	// s := strconv.Itoa(rd)
	// to := []string{user.Email}
	// cc := []string{user.Email}
	// subject := "Email verification"
	// message := "You can also enter this verification code :" + s

	//sendMail(to, cc, subject, message)
	//user.OTP = rd
	// hashPass, eror := PasswordHashing(user.Password)
	// if eror != nil {
	// 	return nil
	// }

	//user.Password = hashPass
	user.Email_verified = false

	sql := `INSERT INTO users (username, name, password, created_date, email, email_verified)
			VALUES ($1, $2, $3, $4, $5, $6)`

	if user.Name == "" || user.Password == "" || user.Username == "" || user.Email == ""{
		return RepoErr
	}
	_, err := r.db.ExecContext(
		ctx, sql, user.Username, user.Name, user.Password, user.Created_date, user.Email, user.Email_verified)
	
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) CheckUsername(ctx context.Context, username string) (bool, error){
	var exists bool

	level.Debug(r.logger).Log("msg", "start run CheckUsername")
	sql := `SELECT EXISTS(SELECT * FROM users WHERE username=$1);`

	err := r.db.QueryRow(sql, username).Scan(&exists)
	if err != nil {
		return false, err
	}

	level.Debug(r.logger).Log("msg", "finish CheckUsername")
	return exists, nil
}