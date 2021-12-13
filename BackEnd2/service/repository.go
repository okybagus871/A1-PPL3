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
	CheckEmail(ctx context.Context, email string) (bool, error)
	ValidateEmail(ctx context.Context, email string, otp uint32) (uint32, error)
	UpdateEmailVerified(ctx context.Context, email string) error
	GetUserByEmail(ctx context.Context, email string) (*datastruct.User, error)
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

func (r *repo) CheckEmail(ctx context.Context, email string) (bool, error){
	var exists bool

	level.Debug(r.logger).Log("msg", "start run CheckEmail")
	sql := `SELECT EXISTS(SELECT * FROM users WHERE email = $1);`

	err := r.db.QueryRow(sql, email).Scan(&exists) 
	if err != nil {
		return false, err
	}
	level.Debug(r.logger).Log("msg", "finish CheckEmail")
	return exists, nil
}

func (r *repo) ValidateEmail(ctx context.Context, email string, otp uint32) (uint32, error){
	var otp_sys uint32
	level.Debug(r.logger).Log("msg", "start run ValidateEmail")

	sql:= `SELECT otp FROM users WHERE email = $1 LIMIT 1;`
	err := r.db.QueryRow(sql, email).Scan(&otp_sys)

	if err != nil {
		return 0, err
	}
	
	level.Debug(r.logger).Log("msg", "finish ValidateEmail")
	return otp_sys, nil
}

func (r *repo) UpdateEmailVerified(ctx context.Context, email string) error {
	level.Debug(r.logger).Log("msg", "start run UpdateEmailVerified")

	email_verified := true
	sql := `UPDATE users SET email_verified = $1 WHERE email = $2;`
	_, err := r.db.ExecContext(ctx, sql, email_verified, email)

	if err != nil {
		return err
	}

	level.Debug(r.logger).Log("msg", "finish UpdateEmailVerified")
	return nil
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*datastruct.User, error) {
	level.Debug(r.logger).Log("msg", "start run GetUserByEmail")

	var user datastruct.User
	sql := `SELECT username, password, email, name, created_date, email_verified FROM users WHERE email = $1;`
	err := r.db.QueryRowContext(ctx, sql, email).Scan(
		&user.Username,
		&user.Password,
		&user.Email,
		&user.Name,
		&user.Created_date,
		&user.Email_verified,
	)

	if err != nil {
		return nil, err
	}


	level.Debug(r.logger).Log("msg", "finish GetUserByEmail")

	return &user, nil
}